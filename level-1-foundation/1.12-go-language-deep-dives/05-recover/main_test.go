package main

import (
	"strings"
	"testing"
)

func TestSafeRun(t *testing.T) {
	err := SafeRun(func() {
		panic("critical crash")
	})
	if err == nil || !strings.Contains(err.Error(), "critical crash") {
		t.Fatalf("expected recovered error, got %v", err)
	}

	err = SafeRun(func() {
		// no panic
	})
	if err != nil {
		t.Errorf("expected nil error on clean run, got %v", err)
	}
}
