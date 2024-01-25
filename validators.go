package go_utils

import (
	"fmt"
	"reflect"
)

// Validate returns any missing fields specified in an object.
func Validate[T any](obj T, fields ...string) ([]string, error) {
	reflectedObject := reflect.ValueOf(&obj).Elem()
	missingFields := make([]string, 0)

	for _, field := range fields {
		value := reflectedObject.FieldByName(field)
		isZero := false

		switch value.Kind() {
		case reflect.Bool:
			isZero = !value.Bool()
		case reflect.Int:
			fallthrough
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			isZero = value.Int() == 0
		case reflect.Uint:
			fallthrough
		case reflect.Uint8:
			fallthrough
		case reflect.Uint16:
			fallthrough
		case reflect.Uint32:
			fallthrough
		case reflect.Uint64:
			fallthrough
		case reflect.Uintptr:
			isZero = value.Uint() == 0
		case reflect.Float32:
			fallthrough
		case reflect.Float64:
			isZero = value.Float() == 0
		case reflect.String:
			isZero = value.String() == ""
		case reflect.Slice:
			fallthrough
		case reflect.Array:
			fallthrough
		case reflect.Map:
			isZero = value.Len() == 0

		default:
			return []string{}, fmt.Errorf("field %s is not supported", field)
		}

		if isZero {
			missingFields = append(missingFields, field)
		}
	}

	return missingFields, nil
}

