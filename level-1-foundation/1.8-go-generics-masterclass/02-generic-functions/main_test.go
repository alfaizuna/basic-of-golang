package main

import "testing"

func TestLast(t *testing.T) {
	ints := []int{1, 2, 99}
	val, ok := Last(ints)
	if !ok || val != 99 {
		t.Errorf("Last(ints) = (%d, %t), want (99, true)", val, ok)
	}

	empty := []string{}
	sVal, sOk := Last(empty)
	if sOk || sVal != "" {
		t.Errorf("Last(empty) = (%q, %t), want ('', false)", sVal, sOk)
	}
}
