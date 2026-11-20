package main

import (
	"math"
	"testing"
)

func TestAddTax(t *testing.T) {
	got := AddTax(200.0)
	want := 220.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("AddTax(200) = %f, want %f", got, want)
	}
}
