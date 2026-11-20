package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestZoneTransferDeepCopy(t *testing.T) {
	primary := NewZone("example.com")
	r1 := NewRecord("example.com", TypeA, "1.1.1.1", 300)
	r2 := NewRecord("www.example.com", TypeCNAME, "example.com", 3600)
	primary.AddRecord(r1)
	primary.AddRecord(r2)

	replica := primary.Transfer()
	if replica == nil {
		t.Fatal("expected non-nil replica from Transfer")
	}
	if replica.Name != primary.Name {
		t.Errorf("expected replica name %q, got %q", primary.Name, replica.Name)
	}
	if replica.RecordCount() != 2 {
		t.Errorf("expected replica count 2, got %d", replica.RecordCount())
	}

	repRec := replica.Lookup("example.com", TypeA)
	if repRec == nil {
		t.Fatal("expected 'example.com' to exist in replica")
	}

	// Pastikan bukan pointer yang sama (deep copy)
	if repRec == r1 {
		t.Errorf("expected deep copy, but replica and primary share the exact same record pointer!")
	}

	// Modifikasi pada replica TIDAK boleh merusak primary
	repRec.Value = "8.8.8.8"
	if r1.Value != "1.1.1.1" {
		t.Errorf("modifying replica affected primary! Expected primary value '1.1.1.1', got %q", r1.Value)
	}

	// Nil safety check
	var nilZone *Zone
	if res := nilZone.Transfer(); res != nil {
		t.Errorf("expected nil result when transferring nil zone")
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
	want1 := "Primary zone count: 2"
	want2 := "Replica zone count: 2"
	want3 := "Independent memory confirmed: true"

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
