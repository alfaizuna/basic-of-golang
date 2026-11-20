package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	f()

	_ = w.Close()
	os.Stdout = stdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

func TestPrintAll(t *testing.T) {
	testData := []any{100, "Go is awesome", true}
	out := captureOutput(func() {
		PrintAll(testData)
	})

	expectedSubstrings := []string{
		"[int] 100",
		"[string] Go is awesome",
		"[bool] true",
	}

	for _, sub := range expectedSubstrings {
		if !strings.Contains(out, sub) {
			t.Errorf("PrintAll output should contain %q, but got:\n%s", sub, out)
		}
	}
}

func TestMainExecution(t *testing.T) {
	out := captureOutput(func() {
		main()
	})

	if !strings.Contains(out, "[int] 100") || !strings.Contains(out, "[string] Go is awesome") || !strings.Contains(out, "[bool] true") {
		t.Errorf("main() output does not match expected, got:\n%s", out)
	}
}
