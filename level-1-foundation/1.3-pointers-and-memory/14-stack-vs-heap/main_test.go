package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSumValues(t *testing.T) {
	got := SumValues(15, 25)
	if got != 40 {
		t.Errorf("expected 40, got %d", got)
	}
}

func TestCreateBuffer(t *testing.T) {
	buf := CreateBuffer(512)
	if buf == nil {
		t.Fatal("expected non-nil pointer from CreateBuffer")
	}
	if len(*buf) != 512 {
		t.Errorf("expected buffer length 512, got %d", len(*buf))
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
	want1 := "Sum: 30"
	want2 := "Buffer len: 1024"

	if !strings.Contains(got, want1) {
		t.Errorf("output does not contain %q, got:\n%s", want1, got)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("output does not contain %q, got:\n%s", want2, got)
	}
}
