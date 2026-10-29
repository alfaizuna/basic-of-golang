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

func TestGenerateSlug(t *testing.T) {
	tests := []struct {
		title       string
		wantSlug    string
		wantErrText string
	}{
		{"  Belajar Go: Dasar #1 (Pemula)!  ", "belajar-go-dasar-1-pemula", ""},
		{"Hello World", "hello-world", ""},
		{"Go 1.26 Released!", "go-126-released", ""},
		{"", "", "input text cannot be empty"},
		{"   \t\n ", "", "input text cannot be empty"},
		{"### ??? !!!", "", "no valid slug characters found"},
	}

	for _, tt := range tests {
		got, err := GenerateSlug(tt.title)
		if tt.wantErrText != "" {
			if err == nil {
				t.Errorf("GenerateSlug(%q) expected error %q, got nil", tt.title, tt.wantErrText)
			} else if err.Error() != tt.wantErrText {
				t.Errorf("GenerateSlug(%q) error = %q; want %q", tt.title, err.Error(), tt.wantErrText)
			}
		} else {
			if err != nil {
				t.Errorf("GenerateSlug(%q) unexpected error: %v", tt.title, err)
			}
			if got != tt.wantSlug {
				t.Errorf("GenerateSlug(%q) = %q; want %q", tt.title, got, tt.wantSlug)
			}
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

	expected := "Input: \"  Belajar Go: Dasar #1 (Pemula)!  \"\nSlug:  \"belajar-go-dasar-1-pemula\"\nError: no valid slug characters found"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
