package gophers

import (
	"errors"
	"fmt"
	"reflect"
)

// ErrFieldNotFound is returned when a requested field does not exist in the struct
var ErrFieldNotFound = errors.New("Fields not found")

// Returns the supplied type with all pointer indirections removed
func derefType(theType reflect.Type) reflect.Type {
	for theType.Kind() == reflect.Pointer {
		theType = theType.Elem()
	}
	return theType
}

// Get names of all fields in the supplied struct type
func GetFieldNames[T any]() []string {
	var theType = derefType(reflect.TypeFor[T]())
	var names = make([]string, theType.NumField())
	for i := 0; i < theType.NumField(); i++ {
		names[i] = theType.Field(i).Name
	}
	return names
}

// Get values of fields with the supplied names in the supplied struct
// Returns an error if any of the supplied names is not a field of the struct;
// values for missing fields are returned as nil
func GetFieldValuesByNames[T any](s T, names []string) (values []any, e error) {
	var theType = derefType(reflect.TypeOf(s))
	var val = reflect.ValueOf(s)
	for val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	var missingFields []string
	for _, name := range names {
		field, ok := theType.FieldByName(name)
		if ok {
			values = append(values, val.FieldByIndex(field.Index).Interface())
		} else {
			values = append(values, nil)
			missingFields = append(missingFields, name)
		}
	}
	if len(missingFields) > 0 {
		e = fmt.Errorf("%w: %v", ErrFieldNotFound, missingFields)
	}
	return values, e
}

// Get names of fields in the supplied struct type marked with tagName:"tagValue"
func GetFieldNamesByTag[T any](tagName, tagValue string) []string {
	var theType = derefType(reflect.TypeFor[T]())
	var names []string
	for field := range theType.Fields() {
		var field = field
		if field.Tag.Get(tagName) == tagValue {
			names = append(names, field.Name)
		}
	}
	return names
}
