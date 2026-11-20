package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestZoneStorage(t *testing.T) {
	z := NewZone("example.com")
	if z == nil {
		t.Fatal("expected non-nil pointer from NewZone")
	}
	if z.Name != "example.com" {
		t.Errorf("expected Zone Name 'example.com', got %q", z.Name)
	}
	if z.records == nil {
		t.Fatal("expected z.records map to be initialized")
	}
	if z.RecordCount() != 0 {
		t.Errorf("expected empty zone count to be 0, got %d", z.RecordCount())
	}

	// Tambah data langsung ke map untuk tes RecordCount
	z.records["example.com"] = []*Record{
		NewRecord("example.com", TypeA, "1.1.1.1", 300),
		NewRecord("example.com", TypeA, "1.0.0.1", 300),
	}
	z.records["mail.example.com"] = []*Record{
		NewRecord("mail.example.com", TypeA, "1.2.3.4", 300),
	}

	if z.RecordCount() != 3 {
		t.Errorf("expected RecordCount 3, got %d", z.RecordCount())
	}

	var nilZone *Zone
	if nilZone.RecordCount() != 0 {
		t.Errorf("expected nil zone count to be 0, got %d", nilZone.RecordCount())
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
	want := "Zone created: example.com (records: 0)"

	if !strings.Contains(got, want) {
		t.Errorf("output does not contain %q, got:\n%s", want, got)
	}
}
