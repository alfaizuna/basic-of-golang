package main

import (
	"reflect"
	"testing"
)

func TestGenerateFibonacci(t *testing.T) {
	ch := GenerateFibonacci(6)
	var got []int
	for v := range ch {
		got = append(got, v)
	}
	expected := []int{0, 1, 1, 2, 3, 5}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
