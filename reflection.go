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
