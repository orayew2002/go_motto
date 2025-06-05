package validator

import (
	"fmt"
	"reflect"
)

func ValidateRequiredFields(s any) error {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct")
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("required")

		if tag == "true" {
			fieldVal := val.Field(i)

			switch fieldVal.Kind() {
			case reflect.String:
				if fieldVal.String() == "" {
					return fmt.Errorf("field %s is required", field.Name)
				}
			case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
				if fieldVal.Int() == 0 {
					return fmt.Errorf("field %s is required", field.Name)
				}
			case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
				if fieldVal.Uint() == 0 {
					return fmt.Errorf("field %s is required", field.Name)
				}
			case reflect.Float64, reflect.Float32:
				if fieldVal.Float() == 0 {
					return fmt.Errorf("field %s is required", field.Name)
				}
			case reflect.Bool:
			case reflect.Ptr, reflect.Interface:
				if fieldVal.IsNil() {
					return fmt.Errorf("field %s is required", field.Name)
				}
			default:
				if reflect.DeepEqual(fieldVal.Interface(), reflect.Zero(fieldVal.Type()).Interface()) {
					return fmt.Errorf("field %s is required", field.Name)
				}
			}
		}
	}

	return nil
}
