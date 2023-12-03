package entity

import (
	"fmt"
	"reflect"
	"time"

	"github.com/bufbuild/connect-go"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1/entityv1alpha1connect"
	"github.com/common-fate/sdk/uid"
	"github.com/fatih/structtag"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
)

type Client struct {
	// Raw returns the underlying GRPC client.
	// It can be used to call methods that we don't have wrappers for yet.
	raw entityv1alpha1connect.EntityServiceClient

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

	rawClient := entityv1alpha1connect.NewEntityServiceClient(opts.HTTPClient, opts.BaseURL, connectOpts...)

	return Client{raw: rawClient, cache: c}
}

// Entities are objects that can be stored in the authz database.
type Entity interface {
	UID() uid.UID
}

// Implementing the parent interface allows an entity to provide parents based on its attributes.
type Parenter interface {
	Parents() []uid.UID
}

func EntityToAPI(e Entity) (*entityv1alpha1.Entity, []*entityv1alpha1.ChildRelation, error) {
	v := reflect.ValueOf(e)
	if v.Kind() == reflect.Pointer {
		v = reflect.Indirect(v)
	}

	if v.Kind() != reflect.Struct {
		return nil, nil, fmt.Errorf("%s is not a struct", v.Type())
	}

	entity := entityv1alpha1.Entity{
		Uid:        e.UID().ToAPI(),
		Attributes: []*entityv1alpha1.Attribute{},
	}

	var children []*entityv1alpha1.ChildRelation

	p, ok := e.(Parenter)
	if ok {
		for _, parent := range p.Parents() {
			err := parent.Valid()
			if err != nil {
				return nil, nil, errors.Wrapf(err, "parsing parents for entity %s", e.UID())
			}

			children = append(children, &entityv1alpha1.ChildRelation{
				Parent: parent.ToAPI(),
				Child:  entity.Uid,
			})
		}
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)

		key := parseTag(string(f.Tag))
		if key == "" {
			// can't parse the tag so continue on
			continue
		}

		if key == "id" {
			// don't put the ID in the attributes
			continue
		}

		// try and parse as an attribute
		attr, err := extractAttr(v.Field(i))
		if err != nil {
			return nil, nil, errors.Wrapf(err, "extracting attributes for entity %s: field %s", e.UID(), key)
		}

		entity.Attributes = append(entity.Attributes, &entityv1alpha1.Attribute{
			Key:   key,
			Value: attr,
		})
	}

	return &entity, children, nil
}

func extractAttr(val reflect.Value) (*entityv1alpha1.Value, error) {
	switch val.Kind() {
	case reflect.Struct:
		// try and parse as an entity reference
		uid, ok := val.Interface().(uid.UID)
		if ok {
			err := uid.Valid()
			if err != nil {
				return nil, err
			}
			return &entityv1alpha1.Value{
				Value: &entityv1alpha1.Value_Entity{
					Entity: uid.ToAPI(),
				},
			}, nil
		}

		// check if it's a timestamp
		if val.Type() == reflect.TypeOf(time.Time{}) {
			timeVal, ok := val.Interface().(time.Time)
			if ok {
				return &entityv1alpha1.Value{
					Value: &entityv1alpha1.Value_Long{
						Long: timeVal.UnixNano() / int64(time.Millisecond),
					},
				}, nil
			}
		}

		// Handle the Record case
		record := &entityv1alpha1.Record{
			Attributes: []*entityv1alpha1.Attribute{},
		}

		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			key := parseTag(string(field.Tag))
			if key == "" {
				continue
			}

			attr, err := extractAttr(val.Field(i))
			if err != nil {
				return nil, err
			}

			record.Attributes = append(record.Attributes, &entityv1alpha1.Attribute{
				Key:   key,
				Value: attr,
			})
		}

		return &entityv1alpha1.Value{
			Value: &entityv1alpha1.Value_Record{
				Record: record,
			},
		}, nil

	case reflect.String:
		return &entityv1alpha1.Value{
			Value: &entityv1alpha1.Value_Str{
				Str: val.String(),
			},
		}, nil

	case reflect.Bool:
		return &entityv1alpha1.Value{
			Value: &entityv1alpha1.Value_Bool{
				Bool: val.Bool(),
			},
		}, nil

	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
		// check if it's a duration
		if val.Type() == reflect.TypeOf(time.Duration(0)) {
			durationVal, ok := val.Interface().(time.Duration)
			if ok {
				return &entityv1alpha1.Value{
					Value: &entityv1alpha1.Value_Long{
						Long: durationVal.Milliseconds(),
					},
				}, nil
			}
		}

		return &entityv1alpha1.Value{
			Value: &entityv1alpha1.Value_Long{
				Long: val.Int(),
			},
		}, nil

	case reflect.Slice:
		// try and parse as set
		set := &entityv1alpha1.Set{
			Values: []*entityv1alpha1.Value{},
		}

		for i := 0; i < val.Len(); i++ {
			attr, err := extractAttr(val.Index(i))
			if err != nil {
				return nil, err
			}
			set.Values = append(set.Values, attr)
		}

		return &entityv1alpha1.Value{
			Value: &entityv1alpha1.Value_Set{
				Set: set,
			},
		}, nil

	case reflect.Pointer:
		// Handle optional type (pointer)
		if val.IsNil() {
			// If the pointer is nil, return a nil value
			return &entityv1alpha1.Value{
				Value: nil,
			}, nil
		}

		// Extract the value from the non-nil pointer
		return extractAttr(val.Elem())
	}

	return nil, fmt.Errorf("extractAttr: unsupported attribute field type: %s", val.Kind())
}

// parseTag parses the 'authz' struct tag. It returns an empty string if the tag cannot be parsed

func parseTag(input string) string {
	tags, err := structtag.Parse(input)
	if err != nil {
		return ""
	}
	authzTag, err := tags.Get("authz")
	if err != nil {
		return ""
	}

	return authzTag.Name
}

func UnmarshalEntity(e *entityv1alpha1.Entity, out Entity) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("output must be a pointer to a struct")
	}

	// Check EntityType matches the UID Type
	if e.Uid == nil || e.Uid.Type != out.UID().Type {
		return fmt.Errorf("entity type mismatch: expected %s, got %s", out.UID().Type, e.Uid.Type)
	}

	v = v.Elem()

	// Handle UID
	if e.Uid != nil && e.Uid.Id != "" {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			tag := parseTag(string(field.Tag))

			if tag == "id" {
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
		tag := parseTag(string(field.Tag))
		keys[tag] = fieldValue
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

	return nil
}

func unmarshalSetValue(setValue *entityv1alpha1.Value, field reflect.Value) error {
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
	case *entityv1alpha1.Value_Entity:
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

	case *entityv1alpha1.Value_Bool:
		if field.Kind() == reflect.Bool {
			field.SetBool(val.Bool)
			return nil
		}

	case *entityv1alpha1.Value_Str:
		if field.Kind() == reflect.String {
			field.SetString(val.Str)
			return nil
		}

	case *entityv1alpha1.Value_Long:
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

	case *entityv1alpha1.Value_Set:
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

	case *entityv1alpha1.Value_Record:
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
		if tag := parseTag(string(field.Tag)); tag == name {
			return v.Field(i), true
		}
	}
	return reflect.Value{}, false
}
