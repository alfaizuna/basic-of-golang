package main

import "testing"

func TestContains(t *testing.T) {
	if !Contains([]int{1, 2, 3}, 2) {
		t.Error("expected true for existing int")
	}
	if Contains([]string{"a", "b"}, "z") {
		t.Error("expected false for missing string")
	}
}
