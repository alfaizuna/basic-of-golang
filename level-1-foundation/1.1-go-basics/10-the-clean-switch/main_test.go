package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestDayType(t *testing.T) {
	tests := []struct {
		day  string
		want string
	}{
		{"senin", "Hari Kerja"},
		{"jumat", "Hari Kerja"},
		{"sabtu", "Akhir Pekan"},
		{"minggu", "Akhir Pekan"},
		{"libur", "Tidak Dikenal"},
	}

	for _, tt := range tests {
		if got := dayType(tt.day); got != tt.want {
			t.Errorf("dayType(%q) = %q; want %q", tt.day, got, tt.want)
		}
	}
}

func TestCategorizeAge(t *testing.T) {
	tests := []struct {
		age  int
		want string
	}{
		{5, "Anak-anak"},
		{12, "Anak-anak"},
		{13, "Remaja"},
		{17, "Remaja"},
		{19, "Remaja"},
		{20, "Dewasa"},
		{40, "Dewasa"},
	}

	for _, tt := range tests {
		if got := categorizeAge(tt.age); got != tt.want {
			t.Errorf("categorizeAge(%d) = %q; want %q", tt.age, got, tt.want)
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

	expected := "Akhir Pekan\nRemaja"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
