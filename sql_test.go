package gophers

import "testing"

func TestSqlValuesAdd(t *testing.T) {
	var sv = &SqlValues{}
	var first = sv.Add("one")
	var second = sv.Add(2)
	if first != "$1" {
		t.Fatalf("Add() = %q, want %q", first, "$1")
	}
	if second != "$2" {
		t.Fatalf("Add() = %q, want %q", second, "$2")
	}
	if len(sv.values) != 2 {
		t.Fatalf("len(values) = %d, want 2", len(sv.values))
	}
	var custom = &SqlValues{key: "?"}
	if got := custom.Add("x"); got != "?1" {
		t.Fatalf("Add() with custom key = %q, want %q", got, "?1")
	}
}
