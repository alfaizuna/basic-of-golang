package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCreateBufferedSlice(t *testing.T) {
	s := createBufferedSlice(2, 5)
	if s == nil {
		t.Fatalf("createBufferedSlice(2, 5) mengembalikan nil; diharapkan slice")
	}
	if len(s) != 2 {
		t.Errorf("len(s) = %d; want 2", len(s))
	}
	if cap(s) != 5 {
		t.Errorf("cap(s) = %d; want 5", cap(s))
	}

	// Cek kasus lain
	s2 := createBufferedSlice(0, 10)
	if len(s2) != 0 || cap(s2) != 10 {
		t.Errorf("createBufferedSlice(0, 10) len/cap = (%d, %d); want (0, 10)", len(s2), cap(s2))
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

	expected := "[100 200]\n2 5"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
