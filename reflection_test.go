package gophers

import (
	"reflect"
	"testing"
)

type testStruct struct {
	Alpha   int    `json:"alpha" validate:"required"`
	Beta    string `json:"beta" validate:"required"`
	Gamma   float64
	Delta   []byte `json:"delta"`
	Epsilon string `validate:"optional"`
}

func TestGetFieldNames(t *testing.T) {
	var want = []string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}
	var got = GetFieldNames[testStruct]()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetFieldNames() = %v, want %v", got, want)
	}
}

func TestGetFieldNamesByTag(t *testing.T) {
	var want = []string{"Alpha"}
	var got = GetFieldNamesByTag[testStruct]("json", "alpha")
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetFieldNamesByTag(%q, %q) = %v, want %v", "json", "alpha", got, want)
	}
}

func TestGetFieldNamesByTagMultiple(t *testing.T) {
	var want = []string{"Alpha", "Beta"}
	var got = GetFieldNamesByTag[testStruct]("validate", "required")
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetFieldNamesByTag(%q, %q) = %v, want %v", "validate", "required", got, want)
	}
}

func TestGetFieldNamesByTagNoMatch(t *testing.T) {
	var want = []string(nil)
	var got = GetFieldNamesByTag[testStruct]("json", "missing")
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetFieldNamesByTag(%q, %q) = %v, want %v", "json", "missing", got, want)
	}
}
