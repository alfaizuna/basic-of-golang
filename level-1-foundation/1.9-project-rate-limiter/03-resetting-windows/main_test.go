package main

import (
	"testing"
	"time"
)

func TestWindowReset(t *testing.T) {
	baseTime := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	currentTime := baseTime
	nowFn := func() time.Time { return currentTime }

	l := NewFixedWindowLimiter(1, 5*time.Second, nowFn)
	if !l.Allow() {
		t.Fatal("1st request should be allowed")
	}
	if l.Allow() {
		t.Fatal("2nd request in same window should be blocked")
	}

	// Majukan waktu 6 detik
	currentTime = currentTime.Add(6 * time.Second)
	if !l.Allow() {
		t.Fatal("request after window expiry should be allowed")
	}
}
