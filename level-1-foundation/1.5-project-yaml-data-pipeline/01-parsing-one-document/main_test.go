package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestParseRow(t *testing.T) {
	yamlInput := `
id: TX-999
item: Mechanical Keyboard
category: electronics
price: 150.50
quantity: 3
`
	row, err := ParseRow([]byte(yamlInput))
	if err != nil {
		t.Fatalf("unexpected error parsing row: %v", err)
	}
	if row == nil {
		t.Fatal("expected non-nil pointer from ParseRow")
	}
	if row.ID != "TX-999" {
		t.Errorf("expected ID 'TX-999', got %q", row.ID)
	}
	if row.Item != "Mechanical Keyboard" {
		t.Errorf("expected Item 'Mechanical Keyboard', got %q", row.Item)
	}
	if row.Category != "electronics" {
		t.Errorf("expected Category 'electronics', got %q", row.Category)
	}
	if row.Price != 150.50 {
		t.Errorf("expected Price 150.50, got %f", row.Price)
	}
	if row.Quantity != 3 {
		t.Errorf("expected Quantity 3, got %d", row.Quantity)
	}
}

func TestParseRowInvalid(t *testing.T) {
	invalidYAML := `
id: [invalid
`
	_, err := ParseRow([]byte(invalidYAML))
	if err == nil {
		t.Errorf("expected error for malformed YAML")
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
	want := "Parsed: TX-101 | Laptop (electronics) | $1200.00 x 2"

	if !strings.Contains(got, want) {
		t.Errorf("output does not contain %q, got:\n%s", want, got)
	}
}
