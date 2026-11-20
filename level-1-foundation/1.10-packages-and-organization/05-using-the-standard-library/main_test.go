package main

import "testing"

func TestTransformInput(t *testing.T) {
	num, upper, err := TransformInput(" 100 items ")
	if err != nil || num != 100 || upper != "100 ITEMS" {
		t.Fatalf("unexpected result: num=%d, upper=%q, err=%v", num, upper, err)
	}

	_, _, err = TransformInput("abc")
	if err == nil {
		t.Error("expected error for non-number input")
	}
}
