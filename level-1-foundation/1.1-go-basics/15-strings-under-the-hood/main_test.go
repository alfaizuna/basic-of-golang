package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCountBytesAndRunes(t *testing.T) {
	tests := []struct {
		text      string
		wantBytes int
		wantRunes int
	}{
		{"Go", 2, 2},
		{"Go 🚀", 7, 4},
		{"世界", 6, 2},
		{"", 0, 0},
		{"Café", 5, 4},
	}

	for _, tt := range tests {
		gotBytes, gotRunes := countBytesAndRunes(tt.text)
		if gotBytes != tt.wantBytes || gotRunes != tt.wantRunes {
			t.Errorf("countBytesAndRunes(%q) = (%d, %d); want (%d, %d)",
				tt.text, gotBytes, gotRunes, tt.wantBytes, tt.wantRunes)
		}
	}
}

func TestNormalizeText(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"  Belajar GOLANG  ", "belajar golang"},
		{"HELLO WORLD", "hello world"},
		{"\t\n  Spasi Banyak \n", "spasi banyak"},
		{"sudah bersih", "sudah bersih"},
	}

	for _, tt := range tests {
		if got := normalizeText(tt.input); got != tt.want {
			t.Errorf("normalizeText(%q) = %q; want %q", tt.input, got, tt.want)
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

	expected := "7 4\nbelajar golang"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
