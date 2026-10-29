package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestIsAllowedSlugChar(t *testing.T) {
	tests := []struct {
		r    rune
		want bool
	}{
		{'a', true},
		{'Z', true},
		{'5', true},
		{' ', true},
		{'-', true},
		{':', false},
		{'!', false},
		{'#', false},
		{'@', false},
		{'🚀', false},
	}

	for _, tt := range tests {
		got := isAllowedSlugChar(tt.r)
		if got != tt.want {
			t.Errorf("isAllowedSlugChar(%q) = %t; want %t", tt.r, got, tt.want)
		}
	}
}

func TestCountInvalidChars(t *testing.T) {
	tests := []struct {
		text string
		want int
	}{
		{"Go 101: Belajar!", 2},
		{"Hello World", 0},
		{"#slug_generator?", 3},
		{"", 0},
	}

	for _, tt := range tests {
		got := countInvalidChars(tt.text)
		if got != tt.want {
			t.Errorf("countInvalidChars(%q) = %d; want %d", tt.text, got, tt.want)
		}
	}
}

func TestCleanText(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"  Go 101: Belajar #1!  ", "go 101 belajar 1"},
		{"Hello World", "hello world"},
		{"### SPECIAL ###", "special"},
		{"", ""},
		{"---", "---"},
	}

	for _, tt := range tests {
		if got := cleanText(tt.input); got != tt.want {
			t.Errorf("cleanText(%q) = %q; want %q", tt.input, got, tt.want)
		}
	}
}

func TestTransformToSlug(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"  Go 101: Belajar #1!  ", "go-101-belajar-1"},
		{"Belajar   Go---Dasar", "belajar-go-dasar"},
		{"---Go 101---", "go-101"},
		{"Special @ Case & Test", "special-case-test"},
		{"", ""},
		{"###", ""},
	}

	for _, tt := range tests {
		if got := transformToSlug(tt.input); got != tt.want {
			t.Errorf("transformToSlug(%q) = %q; want %q", tt.input, got, tt.want)
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

	expected := "Input: \"  Go 101: Belajar #1!  \"\nSlug:  \"go-101-belajar-1\""
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
