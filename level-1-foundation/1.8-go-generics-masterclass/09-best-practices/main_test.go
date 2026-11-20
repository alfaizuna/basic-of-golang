package main

import (
	"reflect"
	"testing"
)

func TestToSlice(t *testing.T) {
	got := ToSlice("a", "b", "c")
	expected := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
