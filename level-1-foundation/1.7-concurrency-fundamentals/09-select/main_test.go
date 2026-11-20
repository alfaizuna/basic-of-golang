package main

import "testing"

func TestFirstResponder(t *testing.T) {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	ch2 <- "B arrived"
	got := FirstResponder(ch1, ch2)
	if got != "B arrived" {
		t.Errorf("expected 'B arrived', got %q", got)
	}
}
