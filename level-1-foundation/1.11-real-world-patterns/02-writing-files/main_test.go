package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveReport(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")

	err := SaveReport(path, "Hello World")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil || string(data) != "Hello World" {
		t.Fatalf("expected 'Hello World', got %q", string(data))
	}
}
