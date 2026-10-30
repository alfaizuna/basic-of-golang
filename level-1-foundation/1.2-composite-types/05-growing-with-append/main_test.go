package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestAddScore(t *testing.T) {
	input := []int{10, 20}
	got := addScore(input, 30)
	expected := []int{10, 20, 30}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("addScore(%v, 30) = %v; want %v", input, got, expected)
	}
}

func TestMainOutput(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	output := strings.TrimSpace(buf.String())

	expected := "Scores: [80 90 100 95]\nLength: 4"
	if output != expected {
		t.Errorf("Output main() tidak sesuai.\nExpected:\n%s\n\nActual:\n%s", expected, output)
	}
}
