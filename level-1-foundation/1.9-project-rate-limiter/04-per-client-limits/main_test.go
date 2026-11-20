package main

import (
	"testing"
	"time"
)

func TestMultiClientLimiter(t *testing.T) {
	l := NewMultiClientLimiter(1, time.Minute, nil)
	if !l.Allow("alice") {
		t.Error("alice 1st request should pass")
	}
	if l.Allow("alice") {
		t.Error("alice 2nd request should be blocked")
	}
	if !l.Allow("bob") {
		t.Error("bob 1st request should pass regardless of alice")
	}
}
