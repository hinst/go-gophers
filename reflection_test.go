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

func TestGetFieldsByTag(t *testing.T) {
	type testCase struct {
		name     string
		tagName  string
		tagValue string
		want     []string
	}
	var cases = []testCase{
		{
			name:     "json tag with value",
			tagName:  "json",
			tagValue: "alpha",
			want:     []string{"Alpha"},
		},
		{
			name:     "multiple fields with same tag value",
			tagName:  "validate",
			tagValue: "required",
			want:     []string{"Alpha", "Beta"},
		},
		{
			name:     "no match returns nil",
			tagName:  "json",
			tagValue: "missing",
			want:     nil,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got = GetFieldsByTag[testStruct](tc.tagName, tc.tagValue)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("GetFieldsByTag(%q, %q) = %v, want %v", tc.tagName, tc.tagValue, got, tc.want)
			}
		})
	}
}
