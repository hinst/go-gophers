package gophers

import "testing"

type exampleStruct struct {
	Cat   int
	Mouse string
}

func TestGetFieldNames(t *testing.T) {
	var names = GetFieldNames[exampleStruct]()
	var ok = names[0] == "Cat" && names[1] == "Mouse" && len(names) == 2
	if !ok {
		t.Error("Wrong names")
	}
}
