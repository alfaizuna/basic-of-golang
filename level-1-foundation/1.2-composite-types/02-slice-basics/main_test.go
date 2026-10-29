package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSumSlice(t *testing.T) {
	tests := []struct {
		numbers []int
		want    int
	}{
		{[]int{15, 25, 35, 45}, 120},
		{[]int{1, 2, 3}, 6},
		{[]int{-5, 5, 10}, 10},
		{[]int{}, 0},
	}

	for _, tt := range tests {
		if got := sumSlice(tt.numbers); got != tt.want {
			t.Errorf("sumSlice(%v) = %d; want %d", tt.numbers, got, tt.want)
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

	expected := "Count: 4\nTotal: 120"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
