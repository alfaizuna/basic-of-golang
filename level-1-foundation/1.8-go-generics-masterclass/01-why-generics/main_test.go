package main

import "testing"

func TestMinFunctions(t *testing.T) {
	if got := MinInt(5, 2); got != 2 {
		t.Errorf("MinInt(5, 2) = %d, want 2", got)
	}
	if got := MinFloat(5.5, 9.2); got != 5.5 {
		t.Errorf("MinFloat(5.5, 9.2) = %f, want 5.5", got)
	}
}
