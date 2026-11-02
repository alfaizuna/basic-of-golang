package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCountWords(t *testing.T) {
	input := []string{"a", "b", "a"}
	res := countWords(input)
	if res["a"] != 2 || res["b"] != 1 {
		t.Errorf("countWords(%v) = %v, want a:2, b:1", input, res)
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

	expected := "Go: 3\nRust: 2\nPython: 1"
	if output != expected {
		t.Errorf("Output main() tidak sesuai.\nExpected:\n%s\n\nActual:\n%s", expected, output)
	}
}
