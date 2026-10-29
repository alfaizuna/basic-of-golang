package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSumEven(t *testing.T) {
	tests := []struct {
		limit int
		want  int
	}{
		{10, 30},
		{5, 6},
		{1, 0},
		{0, 0},
	}

	for _, tt := range tests {
		if got := sumEven(tt.limit); got != tt.want {
			t.Errorf("sumEven(%d) = %d; want %d", tt.limit, got, tt.want)
		}
	}
}

func TestCountUntil(t *testing.T) {
	tests := []struct {
		target int
		want   int
	}{
		{10, 4},
		{9, 3},
		{3, 1},
		{0, 1},
	}

	for _, tt := range tests {
		if got := countUntil(tt.target); got != tt.want {
			t.Errorf("countUntil(%d) = %d; want %d", tt.target, got, tt.want)
		}
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

	expected := "30\n4"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
