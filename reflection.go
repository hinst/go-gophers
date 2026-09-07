package gophers

import "reflect"

// Get names of all fields in the supplied struct type
func GetFieldNames[T any]() []string {
	var theType = reflect.TypeFor[T]()
	var names = make([]string, theType.NumField())
	for i := 0; i < theType.NumField(); i++ {
		names[i] = theType.Field(i).Name
	}
	return names
}

// Get values of fields with the supplied names in the supplied struct
func GetFieldValuesByNames[T any](s T, names []string) []any {
	var theType = reflect.TypeOf(s)
	var values []any
	val := reflect.ValueOf(s)
	for _, name := range names {
		field, ok := theType.FieldByName(name)
		if !ok {
			continue
		}
		values = append(values, val.FieldByIndex(field.Index).Interface())
	}
	return values
}

// Get names of fields in the supplied struct type marked with tagName:"tagValue"
func GetFieldNamesByTag[T any](tagName, tagValue string) []string {
	var theType = reflect.TypeFor[T]()
	var names []string
	for field := range theType.Fields() {
		var field = field
		if field.Tag.Get(tagName) == tagValue {
			names = append(names, field.Name)
		}
	}
	return names
}
