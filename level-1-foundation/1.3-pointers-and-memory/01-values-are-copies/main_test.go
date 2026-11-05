package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestResetCountDoesNotMutateCaller(t *testing.T) {
	val := 50
	resetCount(val)
	if val != 50 {
		t.Errorf("Variabel pemanggil ikut berubah! val = %d; want 50", val)
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

	expected := "Di dalam resetCount: 0\nDi dalam main: 100"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
