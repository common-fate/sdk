package entity

import (
	"fmt"
	"reflect"
	"time"

	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/common-fate/sdk/uid"
)

func Unmarshal(e *entityv1alpha1.Entity, out Entity) error {
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
