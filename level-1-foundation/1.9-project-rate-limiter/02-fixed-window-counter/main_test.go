package main

import (
	"testing"
	"time"
)

func TestFixedWindowLimit(t *testing.T) {
	l := NewFixedWindowLimiter(2, time.Minute, nil)
	if !l.Allow() {
		t.Error("1st call should be allowed")
	}
	if !l.Allow() {
		t.Error("2nd call should be allowed")
	}
	if l.Allow() {
		t.Error("3rd call should be rejected")
	}
}
