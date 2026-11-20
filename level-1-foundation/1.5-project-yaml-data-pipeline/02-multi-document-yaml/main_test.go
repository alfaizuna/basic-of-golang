package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestParseMultiDoc(t *testing.T) {
	input := `
id: 1
item: Alpha
category: cat-a
price: 10
quantity: 1
---
id: 2
item: Beta
category: cat-b
price: 20
quantity: 2
`
	rows, err := ParseMultiDoc(input)
	if err != nil {
		t.Fatalf("unexpected error parsing multi-doc: %v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(rows))
	}
	if rows[0].Item != "Alpha" || rows[1].Item != "Beta" {
		t.Errorf("unexpected rows content: %+v", rows)
	}
}

func TestParseMultiDocEmpty(t *testing.T) {
	rows, err := ParseMultiDoc("")
	if err != nil {
		t.Fatalf("unexpected error for empty yaml: %v", err)
	}
	if len(rows) != 0 {
		t.Errorf("expected 0 rows for empty input, got %d", len(rows))
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
	wantHeader := "Loaded 3 rows from multi-document YAML"
	wantRow1 := "Row 1: TX-101 (Laptop)"
	wantRow2 := "Row 2: TX-102 (Desk)"
	wantRow3 := "Row 3: TX-103 (Monitor)"

	if !strings.Contains(got, wantHeader) {
		t.Errorf("output missing %q, got:\n%s", wantHeader, got)
	}
	if !strings.Contains(got, wantRow1) {
		t.Errorf("output missing %q, got:\n%s", wantRow1, got)
	}
	if !strings.Contains(got, wantRow2) {
		t.Errorf("output missing %q, got:\n%s", wantRow2, got)
	}
	if !strings.Contains(got, wantRow3) {
		t.Errorf("output missing %q, got:\n%s", wantRow3, got)
	}
}
