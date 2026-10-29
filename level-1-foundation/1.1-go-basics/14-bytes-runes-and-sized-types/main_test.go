package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestToByte(t *testing.T) {
	tests := []struct {
		r      rune
		wantB  byte
		wantOk bool
	}{
		{'A', 65, true},
		{'z', 122, true},
		{255, 255, true},
		{'世', 0, false},
		{'🚀', 0, false},
		{256, 0, false},
	}

	for _, tt := range tests {
		gotB, gotOk := toByte(tt.r)
		if gotB != tt.wantB || gotOk != tt.wantOk {
			t.Errorf("toByte(%c) = (%d, %t); want (%d, %t)", tt.r, gotB, gotOk, tt.wantB, tt.wantOk)
		}
	}
}

func TestAddInt8(t *testing.T) {
	tests := []struct {
		a    int8
		b    int8
		want int8
	}{
		{10, 20, 30},
		{127, 1, -128},
		{-128, -1, 127},
	}

	for _, tt := range tests {
		if got := addInt8(tt.a, tt.b); got != tt.want {
			t.Errorf("addInt8(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
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

	expected := "65 true\n0 false\n-128"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
