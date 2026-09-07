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

// Get names of fields in the supplied struct type marked with tagName:"tagValue"
func GetFieldsByTag[T any](tagName, tagValue string) []string {
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
