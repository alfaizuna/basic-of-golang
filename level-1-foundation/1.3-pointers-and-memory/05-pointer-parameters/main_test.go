package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestDoubleFunction(t *testing.T) {
	val := 15
	double(&val)
	if val != 30 {
		t.Errorf("double(&val) gagal mengubah nilai asli: dapat %d, diharapkan 30", val)
	}
}

func TestDoubleNilPointer(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("double(nil) menyebabkan panic: %v", r)
		}
	}()

	double(nil)
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

	expected := "Score setelah didouble: 50"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
