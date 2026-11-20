package main

import (
	"testing"
	"time"
)

func TestCleanup(t *testing.T) {
	cur := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	l := NewLimiterWithCleanup(func() time.Time { return cur })

	l.Touch("old")
	cur = cur.Add(10 * time.Minute)
	l.Touch("new")

	removed := l.Cleanup(5 * time.Minute)
	if removed != 1 {
		t.Errorf("expected 1 removed, got %d", removed)
	}
	if _, exists := l.clients["old"]; exists {
		t.Error("old client should have been removed")
	}
	if _, exists := l.clients["new"]; !exists {
		t.Error("new client should remain")
	}
}
