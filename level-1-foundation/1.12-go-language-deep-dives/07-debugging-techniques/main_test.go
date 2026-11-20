package main

import "testing"

func TestSafeFirst(t *testing.T) {
	v, ok := SafeFirst(nil)
	if ok || v != 0 {
		t.Errorf("nil slice should return (0, false)")
	}

	v, ok = SafeFirst([]int{99})
	if !ok || v != 99 {
		t.Errorf("got (%d, %t), want (99, true)", v, ok)
	}
}
