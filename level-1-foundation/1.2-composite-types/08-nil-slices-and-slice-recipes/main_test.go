package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestRemoveAt(t *testing.T) {
	input := []string{"apple", "banana", "cherry"}
	got := removeAt(input, 1)
	expected := []string{"apple", "cherry"}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("removeAt(%v, 1) = %v; want %v", input, got, expected)
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

	expected := "Languages: [Go Python Ruby]\nIs Nil: true"
	if output != expected {
		t.Errorf("Output main() tidak sesuai.\nExpected:\n%s\n\nActual:\n%s", expected, output)
	}
}
