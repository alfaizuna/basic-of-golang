package main

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	expected := []int{2, 4, 6}

	if !reflect.DeepEqual(evens, expected) {
		t.Errorf("expected %v, got %v", expected, evens)
	}
}
