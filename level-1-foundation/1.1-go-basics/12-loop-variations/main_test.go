package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCountVowels(t *testing.T) {
	tests := []struct {
		text string
		want int
	}{
		{"golang", 2},
		{"level up", 3},
		{"rhythm", 0},
		{"aeiou", 5},
	}

	for _, tt := range tests {
		if got := countVowels(tt.text); got != tt.want {
			t.Errorf("countVowels(%q) = %d; want %d", tt.text, got, tt.want)
		}
	}
}

func TestFindTarget(t *testing.T) {
	r, c := findTarget(5, 7)
	if r != 2 || c != 5 {
		t.Errorf("findTarget(5, 7) = (%d, %d); want (2, 5)", r, c)
	}

	r2, c2 := findTarget(3, 4)
	if r2 != 1 || c2 != 3 {
		t.Errorf("findTarget(3, 4) = (%d, %d); want (1, 3)", r2, c2)
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

	expected := "2\n2 5"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
