package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestNewCounter(t *testing.T) {
	c := NewCounter(42)
	if c == nil {
		t.Fatal("expected non-nil pointer from NewCounter")
	}
	if *c != 42 {
		t.Errorf("expected *c to be 42, got %d", *c)
	}

	// Pastikan dua pemanggilan menghasilkan alamat memori yang independen
	c2 := NewCounter(99)
	if c == c2 {
		t.Errorf("expected distinct memory addresses for different calls")
	}
	if *c2 != 99 {
		t.Errorf("expected *c2 to be 99, got %d", *c2)
	}
}

func TestNewProfile(t *testing.T) {
	p := NewProfile("alfaizuna", "alfa@example.com")
	if p == nil {
		t.Fatal("expected non-nil pointer from NewProfile")
	}
	if p.Username != "alfaizuna" {
		t.Errorf("expected Username 'alfaizuna', got %s", p.Username)
	}
	if p.Email != "alfa@example.com" {
		t.Errorf("expected Email 'alfa@example.com', got %s", p.Email)
	}
}

func TestMainOutput(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)

	got := strings.TrimSpace(buf.String())
	want1 := "Counter: 10"
	want2 := "Profile: budi (budi@example.com)"

	if !strings.Contains(got, want1) {
		t.Errorf("output does not contain %q, got:\n%s", want1, got)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("output does not contain %q, got:\n%s", want2, got)
	}
}
