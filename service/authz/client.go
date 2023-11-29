package authz

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1/authzv1alpha1connect"
	"github.com/common-fate/sdk/service/authz/uid"
	"github.com/fatih/structtag"
	"github.com/patrickmn/go-cache"
)

type Client struct {
	// Raw returns the underlying GRPC client.
	// It can be used to call methods that we don't have wrappers for yet.
	raw authzv1alpha1connect.AuthzServiceClient

	// A cache of recently seen entities. Used for retrieving long-lived
	// attributes such as entity names.
	cache *cache.Cache
}

type Opts struct {
	HTTPClient    connect.HTTPClient
	BaseURL       string
	ClientOptions []connect.ClientOption
	// AttrCacheExpiration defaults to 24h if not provided
	AttrCacheExpiration time.Duration
	// AttrCleanupInterval defaults to 48h if not provided
	AttrCleanupInterval time.Duration
}

func NewClient(opts Opts) Client {
	// client uses GRPC by default as the authz service does not support Buf Connect.
	connectOpts := []connect.ClientOption{connect.WithGRPC()}
	connectOpts = append(connectOpts, opts.ClientOptions...)

	if opts.AttrCleanupInterval == 0 {
		opts.AttrCleanupInterval = 48 * time.Hour
	}

	if opts.AttrCacheExpiration == 0 {
		opts.AttrCacheExpiration = 24 * time.Hour
	}

	c := cache.New(opts.AttrCacheExpiration, opts.AttrCleanupInterval)

	rawClient := authzv1alpha1connect.NewAuthzServiceClient(opts.HTTPClient, opts.BaseURL, connectOpts...)

	return Client{raw: rawClient, cache: c}
}

// Entities are objects that can be stored in the authz database.
type Entity interface {
	EntityType() string
}

// Implementing the parent interface allows an entity to provide parents based on its attributes.
type Parenter interface {
	Parents() []uid.UID
}

func transformToEntity(e Entity) (*authzv1alpha1.Entity, error) {
	v := reflect.ValueOf(e)
	if v.Kind() == reflect.Pointer {
		v = reflect.Indirect(v)
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%s is not a struct", v.Type())
	}

	entity := authzv1alpha1.Entity{
		Attributes: []*authzv1alpha1.Attribute{},
		Parents:    []*authzv1alpha1.UID{},
	}

	p, ok := e.(Parenter)
	if ok {
		for _, parent := range p.Parents() {
			entity.Parents = append(entity.Parents, parent.ToAPI())
		}
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)

		t, err := parseTag(string(f.Tag))
		if err != nil {
			return nil, err
		}

		if t.Name == "id" {
			switch val := v.Field(i); val.Kind() {
			case reflect.String:
				entity.Uid = &authzv1alpha1.UID{
					Type: e.EntityType(),
					Id:   val.String(),
				}

			default:
				return nil, errors.New("unsupported ID field type, only string IDs are currently supported")
			}
			continue
		}

		// try and parse as a parent
		if t.ParentType != "" {
			switch val := v.Field(i); val.Kind() {
			case reflect.String:
				entity.Parents = append(entity.Parents, &authzv1alpha1.UID{
					Type: t.ParentType,
					Id:   val.String(),
				})

			case reflect.Slice:
				slice, ok := val.Interface().([]string)
				if !ok {
					return nil, errors.New("invalid slice: unsupported parent field type, only strings and string slice IDs are currently supported")
				}

				for _, s := range slice {
					entity.Parents = append(entity.Parents, &authzv1alpha1.UID{
						Type: t.ParentType,
						Id:   s,
					})
				}

			default:
				return nil, errors.New("unsupported parent field type, only strings and string slice IDs are currently supported")
			}
			continue
		}

		// try and pass as a generic parent
		if t.HasParent {
			val := v.Field(i)
			slice, ok := val.Interface().(*authzv1alpha1.UID)
			if ok {
				entity.Parents = append(entity.Parents, slice)
				continue
			}

			switch val.Kind() {
			case reflect.Slice:
				slice, ok := val.Interface().([]*authzv1alpha1.UID)
				if !ok {
					return nil, errors.New("invalid slice: unsupported parent field type, []*authzv1alpha1.UID slices are currently supported")
				}

				entity.Parents = append(entity.Parents, slice...)

			default:
				return nil, errors.New("unsupported parent field type, only []*authzv1alpha1.UID slices are currently supported. Otherwise, specify a particular parent entity type with parent=<type>")
			}
			continue
		}

		// try and parse as an attribute
		attr, err := extractAttr(v.Field(i))
		if err != nil {
			return nil, err
		}

		entity.Attributes = append(entity.Attributes, &authzv1alpha1.Attribute{
			Key:   t.Name,
			Value: attr,
		})
	}

	return &entity, nil
}

func extractAttr(val reflect.Value) (*authzv1alpha1.Value, error) {
	switch val.Kind() {
	case reflect.Struct:
		// try and parse as an entity reference
		uid, ok := val.Interface().(uid.UID)
		if ok {
			return &authzv1alpha1.Value{
				Value: &authzv1alpha1.Value_Entity{
					Entity: uid.ToAPI(),
				},
			}, nil
		}

		// check if it's a timestamp
		if val.Type() == reflect.TypeOf(time.Time{}) {
			timeVal, ok := val.Interface().(time.Time)
			if ok {
				return &authzv1alpha1.Value{
					Value: &authzv1alpha1.Value_Long{
						Long: timeVal.UnixNano() / int64(time.Millisecond),
					},
				}, nil
			}
		}

		// Handle the Record case
		record := &authzv1alpha1.Record{
			Attributes: []*authzv1alpha1.Attribute{},
		}

		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			attrTag, err := parseTag(string(field.Tag))
			if err != nil {
				return nil, err
			}

			attr, err := extractAttr(val.Field(i))
			if err != nil {
				return nil, err
			}

			record.Attributes = append(record.Attributes, &authzv1alpha1.Attribute{
				Key:   attrTag.Name,
				Value: attr,
			})
		}

		return &authzv1alpha1.Value{
			Value: &authzv1alpha1.Value_Record{
				Record: record,
			},
		}, nil

	case reflect.String:
		return &authzv1alpha1.Value{
			Value: &authzv1alpha1.Value_Str{
				Str: val.String(),
			},
		}, nil

	case reflect.Bool:
		return &authzv1alpha1.Value{
			Value: &authzv1alpha1.Value_Bool{
				Bool: val.Bool(),
			},
		}, nil

	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
		// check if it's a duration
		if val.Type() == reflect.TypeOf(time.Duration(0)) {
			durationVal, ok := val.Interface().(time.Duration)
			if ok {
				return &authzv1alpha1.Value{
					Value: &authzv1alpha1.Value_Long{
						Long: durationVal.Milliseconds(),
					},
				}, nil
			}
		}

		return &authzv1alpha1.Value{
			Value: &authzv1alpha1.Value_Long{
				Long: val.Int(),
			},
		}, nil

	case reflect.Slice:
		// try and parse as set
		set := &authzv1alpha1.Set{
			Values: []*authzv1alpha1.Value{},
		}

		for i := 0; i < val.Len(); i++ {
			attr, err := extractAttr(val.Index(i))
			if err != nil {
				return nil, err
			}
			set.Values = append(set.Values, attr)
		}

		return &authzv1alpha1.Value{
			Value: &authzv1alpha1.Value_Set{
				Set: set,
			},
		}, nil

	case reflect.Pointer:
		// Handle optional type (pointer)
		if val.IsNil() {
			// If the pointer is nil, return a nil value
			return &authzv1alpha1.Value{
				Value: nil,
			}, nil
		}

		// Extract the value from the non-nil pointer
		return extractAttr(val.Elem())
	}

	return nil, fmt.Errorf("extractAttr: unsupported attribute field type: %s", val.Kind())
}

type Tag struct {
	// Name of the tag, e.g. in `authz:"id"` it is "id"
	Name string

	// the "parent" key e.g. in  `authz:"groups,parent=Group"` it is "Group"
	ParentType string

	// Parent is true if the struct tag is like `authz:"resources,parent"`
	HasParent bool
}

func parseTag(input string) (Tag, error) {
	tags, err := structtag.Parse(input)
	if err != nil {
		return Tag{}, err
	}
	authzTag, err := tags.Get("authz")
	if err != nil {
		return Tag{}, err
	}

	t := Tag{
		Name:       authzTag.Name,
		ParentType: extractParentType(authzTag),
		HasParent:  strings.Contains(input, "parent"),
	}

	return t, nil
}

// extractParentType extracts the parent option from the struct tag.
// e.g. if the tag is `authz:"groups,parent=Group"` the field will be "Group".
func extractParentType(t *structtag.Tag) string {
	for _, opt := range t.Options {
		// split "type=User" into ["type", "User"]
		splits := strings.Split(opt, "=")
		if len(splits) < 2 {
			continue
		}
		if splits[0] == "parent" {
			return splits[1]
		}
	}

	return ""
}

func UnmarshalEntity(e *authzv1alpha1.Entity, out Entity) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("output must be a pointer to a struct")
	}

	// Check EntityType matches the UID Type
	if e.Uid == nil || e.Uid.Type != out.EntityType() {
		return fmt.Errorf("entity type mismatch: expected %s, got %s", out.EntityType(), e.Uid.Type)
	}

	v = v.Elem()

	// Handle UID
	if e.Uid != nil && e.Uid.Id != "" {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			tag, _ := parseTag(string(field.Tag))

			if tag.Name == "id" {
				v.Field(i).SetString(e.Uid.Id)
				break
			}
		}
	}

	// build a map of tag keys to struct fields
	keys := map[string]reflect.Value{}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldValue := v.Field(i)
		tag, _ := parseTag(string(field.Tag))
		keys[tag.Name] = fieldValue
	}

	// Handle Attributes
	for _, attr := range e.Attributes {
		// find the corresponding field
		field, ok := keys[attr.Key]
		if !ok {
			// no field mapped in the struct
			continue
		}

		if err := unmarshalSetValue(attr.Value, field); err != nil {
			return err
		}
	}

	// Handle Parents
	// Assuming Parents are mapped to fields marked with `parent` tag
	for _, parent := range e.Parents {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			tag, _ := parseTag(string(field.Tag))

			if tag.ParentType != "" && parent.Type == tag.ParentType {
				fieldValue := v.Field(i)
				switch fieldValue.Kind() {
				case reflect.String:
					fieldValue.SetString(parent.Id)
				case reflect.Slice:
					// Check if the slice is of type string
					if fieldValue.Type().Elem().Kind() == reflect.String {
						// Append the new string to the slice
						updatedSlice := reflect.Append(fieldValue, reflect.ValueOf(parent.Id))
						fieldValue.Set(updatedSlice)
					}
				}
				break
			}
		}
	}

	return nil
}

func unmarshalSetValue(setValue *authzv1alpha1.Value, field reflect.Value) error {
	if setValue == nil {
		return nil
	}

	// Check if the field is a pointer
	if field.Kind() == reflect.Pointer && setValue != nil && setValue.Value != nil {
		// If it's a pointer, create a new instance and set the value
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		return unmarshalSetValue(setValue, field.Elem())
	}

	switch val := setValue.Value.(type) {
	case *authzv1alpha1.Value_Entity:
		_, ok := field.Interface().(uid.UID)
		if ok {
			field.Set(reflect.ValueOf(uid.UID{
				Type: val.Entity.Type,
				ID:   val.Entity.Id,
			}))
			return nil
		}

		_, ok = field.Interface().(*uid.UID)
		if ok {
			field.Set(reflect.ValueOf(&uid.UID{
				Type: val.Entity.Type,
				ID:   val.Entity.Id,
			}))
			return nil
		}

	case *authzv1alpha1.Value_Bool:
		if field.Kind() == reflect.Bool {
			field.SetBool(val.Bool)
			return nil
		}

	case *authzv1alpha1.Value_Str:
		if field.Kind() == reflect.String {
			field.SetString(val.Str)
			return nil
		}

	case *authzv1alpha1.Value_Long:
		// handle time.Duration
		if field.Type() == reflect.TypeOf(time.Duration(0)) {
			durationVal := time.Duration(val.Long) * time.Millisecond
			field.Set(reflect.ValueOf(durationVal))
			return nil
		}

		kind := field.Kind()
		if kind == reflect.Int ||
			kind == reflect.Int16 ||
			kind == reflect.Int32 ||
			kind == reflect.Int8 ||
			kind == reflect.Int64 {
			field.SetInt(val.Long)
			return nil
		}

		// handle time.Time
		if field.Type() == reflect.TypeOf(time.Time{}) {
			timeVal := time.Unix(0, val.Long*int64(time.Millisecond))
			field.Set(reflect.ValueOf(timeVal))
			return nil
		}

	case *authzv1alpha1.Value_Set:
		// Handle the Set case
		switch field.Kind() {
		case reflect.Slice:
			sliceType := field.Type().Elem()
			slice := reflect.MakeSlice(field.Type(), 0, len(val.Set.Values))
			for _, nestedSetValue := range val.Set.Values {
				nestedSetField := reflect.New(sliceType).Elem()
				if err := unmarshalSetValue(nestedSetValue, nestedSetField); err != nil {
					return err
				}
				slice = reflect.Append(slice, nestedSetField)
			}
			field.Set(slice)
			return nil
		}

	case *authzv1alpha1.Value_Record:
		// Handle the Record case
		switch field.Kind() {
		case reflect.Struct:
			for _, attr := range val.Record.Attributes {
				// find the corresponding field
				attrField, ok := getFieldByName(field, attr.Key)
				if !ok {
					// no field mapped in the struct
					continue
				}

				// recursively unmarshal the attribute value into the field
				if err := unmarshalSetValue(attr.Value, attrField); err != nil {
					return err
				}
			}

			return nil
		}
	case nil:
		// Handle the case where the value is nil
		return nil

	default:
		return fmt.Errorf("unsupported attribute field type: %T", val)
	}
	return fmt.Errorf("unhandled attribute field: %+v", setValue)
}

// getFieldByName finds a field in a struct by its name
func getFieldByName(v reflect.Value, name string) (reflect.Value, bool) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if tag, _ := parseTag(string(field.Tag)); tag.Name == name {
			return v.Field(i), true
		}
	}
	return reflect.Value{}, false
}
