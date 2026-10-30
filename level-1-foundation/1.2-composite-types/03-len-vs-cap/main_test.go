package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestGetLenAndCap(t *testing.T) {
	data := []int{10, 20, 30, 40, 50}
	part := data[1:3]

	gotLen, gotCap := getLenAndCap(part)
	if gotLen != 2 {
		t.Errorf("getLenAndCap(part) len = %d; want 2", gotLen)
	}
	if gotCap != 4 {
		t.Errorf("getLenAndCap(part) cap = %d; want 4", gotCap)
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

	expected := "2 4"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
