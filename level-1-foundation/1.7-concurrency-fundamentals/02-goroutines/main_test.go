package main

import (
	"testing"
	"time"
)

func TestLaunchWorker(t *testing.T) {
	out := make(chan string, 1)
	LaunchWorker(42, out)

	select {
	case msg := <-out:
		if msg != "Worker 42 done" {
			t.Errorf("expected 'Worker 42 done', got %q", msg)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("timeout waiting for worker goroutine")
	}
}
