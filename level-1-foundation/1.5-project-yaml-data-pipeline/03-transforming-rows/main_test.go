package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCalculateTotals(t *testing.T) {
	input := []Row{
		{Price: 100.5, Quantity: 2},
		{Price: 50.0, Quantity: 3},
	}
	out := CalculateTotals(input)
	if len(out) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(out))
	}
	if out[0].Total != 201.0 {
		t.Errorf("expected Total 201.0, got %f", out[0].Total)
	}
	if out[1].Total != 150.0 {
		t.Errorf("expected Total 150.0, got %f", out[1].Total)
	}
}

func TestNormalizeCategories(t *testing.T) {
	input := []Row{
		{Category: "  Electronics "},
		{Category: "FURNITURE"},
	}
	out := NormalizeCategories(input)
	if out[0].Category != "electronics" {
		t.Errorf("expected 'electronics', got %q", out[0].Category)
	}
	if out[1].Category != "furniture" {
		t.Errorf("expected 'furniture', got %q", out[1].Category)
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
	want1 := "Transformed: TX-101 | Laptop | electronics | Total: $2400.00"
	want2 := "Transformed: TX-102 | Desk | furniture | Total: $450.00"

	if !strings.Contains(got, want1) {
		t.Errorf("output does not contain %q, got:\n%s", want1, got)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("output does not contain %q, got:\n%s", want2, got)
	}
}
