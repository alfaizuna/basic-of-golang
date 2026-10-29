package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCalculate(t *testing.T) {
	s1, d1 := calculate(20, 5)
	if s1 != 25 || d1 != 15 {
		t.Errorf("calculate(20, 5) = (%d, %d); want (25, 15)", s1, d1)
	}

	s2, d2 := calculate(10, 15)
	if s2 != 25 || d2 != -5 {
		t.Errorf("calculate(10, 15) = (%d, %d); want (25, -5)", s2, d2)
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

	expected := "25 15\n60"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
