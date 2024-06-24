package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
)

func GetBodyInJson(r *http.Request, bodyType interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, bodyType)
	if err != nil {
		return err
	}

	return nil
}

func GetNullableValue(value *string) interface{} {
	if value == nil {
		return nil
	}
	return *value
}

// StructToMap converts a struct instance to a map[string]interface{}.
func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(data)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		key := strings.ToLower(string(field.Name[0])) + field.Name[1:]
		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				continue // Omit nil pointers
			}
			value = value.Elem() // Dereference pointer
		}
		if isEmpty(value) {
			continue // Omit empty values
		}
		result[key] = value.Interface()
	}

	return result
}

// isEmpty checks if a value is empty.
func isEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.String:
		return v.Len() == 0
	case reflect.Array, reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.Struct:
		zero := reflect.Zero(v.Type())
		return reflect.DeepEqual(v.Interface(), zero.Interface())
	}
	return false
}
