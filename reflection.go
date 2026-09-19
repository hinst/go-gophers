package gophers

import (
	"errors"
	"fmt"
	"reflect"
)

// ErrFieldNotFound is returned when a requested field does not exist in the struct
var ErrFieldNotFound = errors.New("Fields not found")

func unwrapPointerType(theType reflect.Type) reflect.Type {
	for theType.Kind() == reflect.Pointer {
		theType = theType.Elem()
	}
	return theType
}

// Get names of all fields in the supplied struct type
func GetFieldNames[T any]() []string {
	var theType = unwrapPointerType(reflect.TypeFor[T]())
	var names = make([]string, theType.NumField())
	for i := 0; i < theType.NumField(); i++ {
		names[i] = theType.Field(i).Name
	}
	return names
}

// Get the value of the field with the supplied name in the supplied struct
// Returns an error if the supplied name is not a field of the struct
func GetFieldValueByName[T any](s T, name string) (value any, e error) {
	var theType = unwrapPointerType(reflect.TypeOf(s))
	var val = reflect.ValueOf(s)
	for val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	field, ok := theType.FieldByName(name)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrFieldNotFound, name)
	}
	return val.FieldByIndex(field.Index).Interface(), nil
}

// Get values of fields with the supplied names in the supplied struct
// Returns an error if any of the supplied names is not a field of the struct;
// values for missing fields are returned as nil
func GetFieldValuesByNames[T any](s T, names []string) (values []any, e error) {
	var missingFields []string
	for _, name := range names {
		value, err := GetFieldValueByName(s, name)
		if err != nil {
			values = append(values, nil)
			missingFields = append(missingFields, name)
		} else {
			values = append(values, value)
		}
	}
	if len(missingFields) > 0 {
		e = fmt.Errorf("%w: %v", ErrFieldNotFound, missingFields)
	}
	return values, e
}

// Get names of fields in the supplied struct type marked with tagName:"tagValue"
func GetFieldNamesByTag[T any](tagName, tagValue string) []string {
	var theType = unwrapPointerType(reflect.TypeFor[T]())
	var names []string
	for field := range theType.Fields() {
		var field = field
		if field.Tag.Get(tagName) == tagValue {
			names = append(names, field.Name)
		}
	}
	return names
}
