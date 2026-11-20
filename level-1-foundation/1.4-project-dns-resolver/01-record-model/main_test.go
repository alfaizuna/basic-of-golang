package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestRecordModel(t *testing.T) {
	if string(TypeA) != "A" {
		t.Errorf("expected TypeA to be 'A', got %q", TypeA)
	}
	if string(TypeCNAME) != "CNAME" {
		t.Errorf("expected TypeCNAME to be 'CNAME', got %q", TypeCNAME)
	}

	rec := NewRecord("example.com", TypeA, "1.1.1.1", 120)
	if rec == nil {
		t.Fatal("expected non-nil pointer from NewRecord")
	}
	if rec.Name != "example.com" {
		t.Errorf("expected Name 'example.com', got %q", rec.Name)
	}
	if rec.Type != TypeA {
		t.Errorf("expected Type TypeA, got %v", rec.Type)
	}
	if rec.Value != "1.1.1.1" {
		t.Errorf("expected Value '1.1.1.1', got %q", rec.Value)
	}
	if rec.TTL != 120 {
		t.Errorf("expected TTL 120, got %d", rec.TTL)
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
	want1 := "Record created: example.com (A -> 93.184.216.34, TTL: 300s)"
	want2 := "Alias created: www.example.com (CNAME -> example.com, TTL: 3600s)"

	if !strings.Contains(got, want1) {
		t.Errorf("output does not contain %q, got:\n%s", want1, got)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("output does not contain %q, got:\n%s", want2, got)
	}
}
