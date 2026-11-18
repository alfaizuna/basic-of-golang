package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestFindUserFound(t *testing.T) {
	u, err := FindUser(1)
	if err != nil {
		t.Fatalf("expected no error for existing ID, got: %v", err)
	}
	if u == nil {
		t.Fatalf("expected non-nil user pointer, got nil")
	}
	if u.Name != "Alice" {
		t.Errorf("expected user name 'Alice', got '%s'", u.Name)
	}
}

func TestFindUserNotFound(t *testing.T) {
	u, err := FindUser(99)
	if err == nil {
		t.Fatalf("expected error for non-existent ID, got nil")
	}
	if u != nil {
		t.Errorf("expected nil user pointer when not found, got %v", u)
	}
	if err.Error() != "user not found" {
		t.Errorf("expected error message 'user not found', got '%s'", err.Error())
	}
}

func TestFormatUserName(t *testing.T) {
	// Test nil guard
	gotNil := FormatUserName(nil)
	if gotNil != "Guest" {
		t.Errorf("expected 'Guest' for nil pointer, got '%s'", gotNil)
	}

	// Test valid pointer
	user := &User{Name: "Bob"}
	gotValid := FormatUserName(user)
	if gotValid != "Bob" {
		t.Errorf("expected 'Bob', got '%s'", gotValid)
	}
}

func TestMainOutput(t *testing.T) {
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	output := strings.TrimSpace(buf.String())

	expectedLines := []string{
		"User: Alice",
		"Error: user not found",
		"Fallback: Guest",
	}

	expected := strings.Join(expectedLines, "\n")
	if output != expected {
		t.Errorf("unexpected output:\nGOT:\n%s\n\nWANT:\n%s", output, expected)
	}
}
