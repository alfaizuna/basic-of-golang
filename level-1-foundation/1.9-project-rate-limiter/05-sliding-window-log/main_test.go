package main

import (
	"testing"
	"time"
)

func TestSlidingWindowLog(t *testing.T) {
	cur := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	s := NewSlidingWindowLog(2, 10*time.Second, func() time.Time { return cur })

	if !s.Allow() || !s.Allow() {
		t.Fatal("first two requests should be allowed")
	}
	if s.Allow() {
		t.Fatal("third request should be blocked")
	}

	// Majukan waktu 11 detik sehingga log lama kedaluwarsa
	cur = cur.Add(11 * time.Second)
	if !s.Allow() {
		t.Fatal("request after window slide should be allowed")
	}
}
