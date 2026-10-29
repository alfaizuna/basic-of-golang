package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestModifyCopy(t *testing.T) {
	input := [3]int{1, 2, 3}
	got := modifyCopy(input)

	if got[0] != 99 {
		t.Errorf("modifyCopy(input)[0] = %d; want 99", got[0])
	}
	if got[1] != 2 || got[2] != 3 {
		t.Errorf("modifyCopy mengubah elemen lain yang tidak seharusnya: %v", got)
	}
	// Pastikan input asli tidak berubah (bukti array adalah value type)
	if input[0] != 1 {
		t.Errorf("Array asli ikut berubah! input[0] = %d; want 1", input[0])
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

	expected := "[10 20 30]\n[99 20 30]"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
