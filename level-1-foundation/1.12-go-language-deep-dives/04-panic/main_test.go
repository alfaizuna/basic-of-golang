package main

import "testing"

func TestRequirePositive(t *testing.T) {
	// Test positive (should not panic)
	RequirePositive(10)

	// Test negative (must panic)
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("expected panic on non-positive value, but got none")
		}
	}()
	RequirePositive(-1)
}
