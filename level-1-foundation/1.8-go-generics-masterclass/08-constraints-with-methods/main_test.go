package main

import "testing"

func TestSumAreas(t *testing.T) {
	squares := []Square{{Side: 3}, {Side: 4}}
	total := SumAreas(squares)
	if total != 25.0 {
		t.Errorf("expected total 25.0, got %f", total)
	}
}
