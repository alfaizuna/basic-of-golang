package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestAddAndLookup(t *testing.T) {
	z := NewZone("example.com")
	r1 := NewRecord("example.com", TypeA, "1.1.1.1", 300)
	r2 := NewRecord("www.example.com", TypeCNAME, "example.com", 3600)

	z.AddRecord(r1)
	z.AddRecord(r2)

	gotA := z.Lookup("example.com", TypeA)
	if gotA != r1 {
		t.Errorf("expected to lookup r1, got %v", gotA)
	}

	gotCNAME := z.Lookup("www.example.com", TypeCNAME)
	if gotCNAME != r2 {
		t.Errorf("expected to lookup r2, got %v", gotCNAME)
	}

	if missing := z.Lookup("unknown.com", TypeA); missing != nil {
		t.Errorf("expected nil for non-existent domain, got %v", missing)
	}

	if wrongType := z.Lookup("example.com", TypeCNAME); wrongType != nil {
		t.Errorf("expected nil for mismatched record type, got %v", wrongType)
	}

	// Nil safety checks
	var nilZone *Zone
	nilZone.AddRecord(r1)
	if res := nilZone.Lookup("example.com", TypeA); res != nil {
		t.Errorf("expected nil lookup on nil zone")
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
	want1 := "Lookup example.com (A): 93.184.216.34"
	want2 := "Lookup notfound.com (A): nil"

	if !strings.Contains(got, want1) {
		t.Errorf("output does not contain %q, got:\n%s", want1, got)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("output does not contain %q, got:\n%s", want2, got)
	}
}
