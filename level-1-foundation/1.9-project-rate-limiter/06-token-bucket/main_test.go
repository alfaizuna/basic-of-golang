package main

import (
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	cur := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	tb := NewTokenBucket(2.0, 0.5, func() time.Time { return cur }) // 0.5 token per detik

	if !tb.Allow() || !tb.Allow() {
		t.Fatal("initial burst should allow 2 tokens")
	}
	if tb.Allow() {
		t.Fatal("bucket should be empty")
	}

	// Majukan 2 detik -> refill 1 token
	cur = cur.Add(2 * time.Second)
	if !tb.Allow() {
		t.Fatal("should have refilled 1 token after 2 seconds")
	}
	if tb.Allow() {
		t.Fatal("bucket should be empty again")
	}
}
