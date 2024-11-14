package entity

import (
	"fmt"
	"reflect"
	"time"

	"github.com/common-fate/sdk/eid"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/fatih/structtag"
	"github.com/pkg/errors"
)

// Entities are objects that can be stored in the authz database.
type Entity interface {
	EID() eid.EID
}

// Implementing the parent interface allows an entity to provide parents based on its attributes.
type Parenter interface {
	Parents() []eid.EID
}

func Marshal(e Entity) (*entityv1alpha1.Entity, []*entityv1alpha1.ChildRelation, error) {
	v := reflect.ValueOf(e)
	if v.Kind() == reflect.Pointer {
		v = reflect.Indirect(v)
	}

	if v.Kind() != reflect.Struct {
		return nil, nil, fmt.Errorf("%s is not a struct", v.Type())
	}

	entity := entityv1alpha1.Entity{
		Eid:        e.EID().ToAPI(),
		Attributes: []*entityv1alpha1.Attribute{},
	}

	var children []*entityv1alpha1.ChildRelation

	p, ok := e.(Parenter)
	if ok {
		for _, parent := range p.Parents() {
			err := parent.Valid()
			if err != nil {
				return nil, nil, errors.Wrapf(err, "parsing parents for entity %s", e.EID())
			}

			children = append(children, &entityv1alpha1.ChildRelation{
				Parent: parent.ToAPI(),
				Child:  entity.Eid,
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
			return nil, nil, errors.Wrapf(err, "extracting attributes for entity %s: field %s", e.EID(), key)
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
		eid, ok := val.Interface().(eid.EID)
		if ok {
			err := eid.Valid()
			if err != nil {
				return nil, err
			}
			return &entityv1alpha1.Value{
				Value: &entityv1alpha1.Value_Entity{
					Entity: eid.ToAPI(),
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

	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
		return &entityv1alpha1.Value{
			Value: &entityv1alpha1.Value_Long{
				Long: int64(val.Uint()),
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
	case reflect.Map:
		if val.Type() != reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf("")) {
			return nil, fmt.Errorf("input is not a map[string]string")
		}

		record := &entityv1alpha1.Record{
			Attributes: []*entityv1alpha1.Attribute{},
		}

		for _, keyVal := range val.MapKeys() {
			key := keyVal.String()
			value := val.MapIndex(keyVal)

			// Assuming value is also a string
			attr := value.String()

			record.Attributes = append(record.Attributes, &entityv1alpha1.Attribute{
				Key: key,
				Value: &entityv1alpha1.Value{
					Value: &entityv1alpha1.Value_Str{
						Str: attr,
					},
				},
			})
		}

		return &entityv1alpha1.Value{
			Value: &entityv1alpha1.Value_Record{
				Record: record,
			},
		}, nil

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
