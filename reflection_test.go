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

func TestGetFieldValuesByNamesAll(t *testing.T) {
	var s = testStruct{
		Alpha:   1,
		Beta:    "two",
		Gamma:   3.5,
		Delta:   []byte{4, 5, 6},
		Epsilon: "five",
	}
	var want = []any{1, "two", 3.5, []byte{4, 5, 6}, "five"}
	var got, err = GetFieldValuesByNames(s, []string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"})
	if err != nil {
		t.Fatalf("GetFieldValuesByNames() returned error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetFieldValuesByNames() = %v, want %v", got, want)
	}
}

func TestGetFieldValuesByNamesSubset(t *testing.T) {
	var s = testStruct{Alpha: 7, Beta: "seven"}
	var want = []any{"seven", 7}
	var got, err = GetFieldValuesByNames(s, []string{"Beta", "Alpha"})
	if err != nil {
		t.Fatalf("GetFieldValuesByNames() returned error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetFieldValuesByNames() = %v, want %v", got, want)
	}
}

func TestGetFieldValuesByNamesMissing(t *testing.T) {
	var s = testStruct{Alpha: 1}
	var want = []any{1, nil, ""}
	var got, err = GetFieldValuesByNames(s, []string{"Alpha", "Missing", "Epsilon"})
	if err == nil {
		t.Fatalf("GetFieldValuesByNames() = %v, want error for missing field", got)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetFieldValuesByNames() = %v, want %v (nil placeholder for missing field)", got, want)
	}
}

func TestGetFieldValuesByNamesNoneFound(t *testing.T) {
	var s = testStruct{Alpha: 1}
	var want = []any{nil}
	var got, err = GetFieldValuesByNames(s, []string{"Missing"})
	if err == nil {
		t.Fatalf("GetFieldValuesByNames() = %v, want error for missing field", got)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetFieldValuesByNames() = %v, want %v (nil placeholder for missing field)", got, want)
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
