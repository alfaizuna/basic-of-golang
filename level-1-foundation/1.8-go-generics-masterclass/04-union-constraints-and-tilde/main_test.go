package main

import "testing"

func TestMaxGenerics(t *testing.T) {
	if got := Max(10, 20); got != 20 {
		t.Errorf("Max(10, 20) = %d, want 20", got)
	}
	s1 := Speed(120.5)
	s2 := Speed(90.0)
	if got := Max(s1, s2); got != s1 {
		t.Errorf("Max(s1, s2) = %v, want %v", got, s1)
	}
}
