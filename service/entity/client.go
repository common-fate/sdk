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
