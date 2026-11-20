package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestUpdateInPlace(t *testing.T) {
	z := NewZone("example.com")
	r := NewRecord("example.com", TypeA, "1.1.1.1", 300)
	z.AddRecord(r)

	// Test UpdateRecord
	ok := z.UpdateRecord("example.com", TypeA, "1.0.0.1", 600)
	if !ok {
		t.Errorf("expected UpdateRecord to return true for existing record")
	}

	if r.Value != "1.0.0.1" {
		t.Errorf("expected original record value to be updated to '1.0.0.1', got %q", r.Value)
	}
	if r.TTL != 600 {
		t.Errorf("expected original record TTL to be updated to 600, got %d", r.TTL)
	}

	// Update non-existing record
	fail := z.UpdateRecord("unknown.com", TypeA, "2.2.2.2", 100)
	if fail {
		t.Errorf("expected UpdateRecord to return false for non-existent record")
	}

	// Direct method tests with nil safety
	var nilRec *Record
	nilRec.UpdateValue("test")
	nilRec.UpdateTTL(50)
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
	want1 := "Before: example.com -> 93.184.216.34 (TTL: 300)"
	want2 := "Update success: true"
	want3 := "After : example.com -> 93.184.216.35 (TTL: 600)"

	if !strings.Contains(got, want1) {
		t.Errorf("output does not contain %q, got:\n%s", want1, got)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("output does not contain %q, got:\n%s", want2, got)
	}
	if !strings.Contains(got, want3) {
		t.Errorf("output does not contain %q, got:\n%s", want3, got)
	}
}
