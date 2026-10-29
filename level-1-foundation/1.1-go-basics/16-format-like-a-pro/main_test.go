package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestFormatProduct(t *testing.T) {
	tests := []struct {
		name  string
		price float64
		stock int
		want  string
	}{
		{"Mechanical Keyboard", 89.9, 15, `Product: "Mechanical Keyboard" | Price: $89.90 | Stock: 15`},
		{"Coffee", 3.5, 0, `Product: "Coffee" | Price: $3.50 | Stock: 0`},
		{"Monitor 4K", 299.999, 5, `Product: "Monitor 4K" | Price: $300.00 | Stock: 5`},
	}

	for _, tt := range tests {
		got := formatProduct(tt.name, tt.price, tt.stock)
		if got != tt.want {
			t.Errorf("formatProduct(%q, %v, %v) = %q; want %q", tt.name, tt.price, tt.stock, got, tt.want)
		}
	}
}

func TestDescribeValue(t *testing.T) {
	tests := []struct {
		val  int
		want string
	}{
		{42, "Type: int | Padded: 0042"},
		{7, "Type: int | Padded: 0007"},
		{12345, "Type: int | Padded: 12345"},
	}

	for _, tt := range tests {
		got := describeValue(tt.val)
		if got != tt.want {
			t.Errorf("describeValue(%v) = %q; want %q", tt.val, got, tt.want)
		}
	}
}

func TestMainOutput(t *testing.T) {
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Gagal membuat pipe: %v", err)
	}
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		outC <- buf.String()
	}()

	main()

	_ = w.Close()
	os.Stdout = oldStdout
	out := <-outC

	expected := "Product: \"Mechanical Keyboard\" | Price: $89.90 | Stock: 15\nType: int | Padded: 0042"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
