package main

import (
	"testing"
	"time"
)

func TestNewLimiter(t *testing.T) {
	fixedTime := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	l := NewLimiter(10, func() time.Time { return fixedTime })

	if l.Limit != 10 {
		t.Errorf("expected limit 10, got %d", l.Limit)
	}
	if !l.Now().Equal(fixedTime) {
		t.Errorf("injected time mismatch")
	}
	if !l.Allow() {
		t.Errorf("expected Allow to return true")
	}
}
