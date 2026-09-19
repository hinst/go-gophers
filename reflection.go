package gophers

import (
	"errors"
	"fmt"
	"reflect"
)

// The requested field does not exist in the struct
var ErrorFieldNotFound = errors.New("Fields not found")

// The requested field is private, therefore we cannot read its value
var ErrorFieldNotExported = errors.New("Field not exported")

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
func GetFieldValueByName[T any](s T, name string) (value any, exception error) {
	var theType = unwrapPointerType(reflect.TypeOf(s))
	var val = reflect.ValueOf(s)
	for val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	field, ok := theType.FieldByName(name)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrorFieldNotFound, name)
	}
	if !field.IsExported() {
		return nil, fmt.Errorf("%w: %s", ErrorFieldNotExported, name)
	}
	return val.FieldByIndex(field.Index).Interface(), nil
}

// Get values of fields with the supplied names.
// Returns error if some of the fields cannot be retrieved.
// Fields that could not be retrieved are returned as nil elements in the output array
func GetFieldValuesByNames[T any](s T, names []string) (values []any, exception error) {
	var errorCount int
	for _, name := range names {
		var value, fieldError = GetFieldValueByName(s, name)
		if fieldError != nil {
			values = append(values, nil)
			errorCount++
			exception = fieldError
		} else {
			values = append(values, value)
		}
	}
	if errorCount > 0 {
		exception = fmt.Errorf("%v of %v fields cannot be retrieved, last error: %w",
			errorCount, len(names), exception)
	}
	return values, exception
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
