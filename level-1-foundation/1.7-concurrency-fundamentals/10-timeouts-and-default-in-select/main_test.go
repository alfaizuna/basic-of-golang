package main

import (
	"testing"
	"time"
)

func TestFetchWithTimeoutSuccess(t *testing.T) {
	ch := make(chan string, 1)
	ch <- "OK"

	val, err := FetchWithTimeout(ch, 100*time.Millisecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "OK" {
		t.Errorf("expected 'OK', got %q", val)
	}
}

func TestFetchWithTimeoutFail(t *testing.T) {
	ch := make(chan string) // Unbuffered, no sender

	_, err := FetchWithTimeout(ch, 20*time.Millisecond)
	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}
}
