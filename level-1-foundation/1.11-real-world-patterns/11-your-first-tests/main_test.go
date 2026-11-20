package main

import "testing"

func TestMultiply(t *testing.T) {
	if got := Multiply(3, 4); got != 12 {
		t.Errorf("Multiply(3, 4) = %d, want 12", got)
	}
}

func TestDivide(t *testing.T) {
	res, err := Divide(10, 2)
	if err != nil || res != 5 {
		t.Errorf("Divide(10, 2) = (%d, %v), want (5, nil)", res, err)
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Fatal("expected error on divide by zero, got nil")
	}
}
