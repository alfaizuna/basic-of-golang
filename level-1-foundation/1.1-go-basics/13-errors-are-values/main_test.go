package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestDivide(t *testing.T) {
	// Test kasus sukses
	res, err := divide(10, 2)
	if err != nil {
		t.Fatalf("divide(10, 2) mengembalikan error: %v, diharapkan nil", err)
	}
	if res != 5 {
		t.Errorf("divide(10, 2) = %v; want 5", res)
	}

	res2, err2 := divide(7, 2)
	if err2 != nil {
		t.Fatalf("divide(7, 2) mengembalikan error: %v, diharapkan nil", err2)
	}
	if res2 != 3.5 {
		t.Errorf("divide(7, 2) = %v; want 3.5", res2)
	}

	// Test kasus error pembagian dengan nol
	resZero, errZero := divide(10, 0)
	if errZero == nil {
		t.Fatalf("divide(10, 0) harus mengembalikan error, tapi dapat nil")
	}
	if errZero.Error() != "cannot divide by zero" {
		t.Errorf("Pesan error divide(10, 0) = %q; want %q", errZero.Error(), "cannot divide by zero")
	}
	if resZero != 0 {
		t.Errorf("divide(10, 0) result = %v; want 0", resZero)
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

	expected := "5\ncannot divide by zero"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
