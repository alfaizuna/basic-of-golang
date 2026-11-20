package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	z := NewZone("example.com")
	z.AddRecord(NewRecord("a.com", TypeCNAME, "b.com", 300))
	z.AddRecord(NewRecord("b.com", TypeCNAME, "a.com", 300))

	_, err := z.Resolve("a.com")
	if err == nil {
		t.Fatal("expected error for circular CNAME chain")
	}
	if !strings.Contains(err.Error(), "cname loop detected") {
		t.Errorf("expected error message to contain 'cname loop detected', got %v", err)
	}
}

func TestMaxDepthExceeded(t *testing.T) {
	z := NewZone("example.com")
	// Buat rantai sepanjang 15 lompatan tanpa loop
	for i := 1; i <= 15; i++ {
		from := fmt.Sprintf("hop-%d.com", i)
		to := fmt.Sprintf("hop-%d.com", i+1)
		z.AddRecord(NewRecord(from, TypeCNAME, to, 300))
	}
	z.AddRecord(NewRecord("hop-16.com", TypeA, "1.2.3.4", 300))

	_, err := z.Resolve("hop-1.com")
	if err == nil {
		t.Fatal("expected error when depth exceeds limit")
	}
	if !strings.Contains(err.Error(), "maximum CNAME depth exceeded") {
		t.Errorf("expected 'maximum CNAME depth exceeded', got %v", err)
	}
}

func TestMainOutput(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)

	got := strings.TrimSpace(buf.String())
	want := "Resolving loop-a: error: cname loop detected"

	if !strings.Contains(got, want) {
		t.Errorf("output does not contain %q, got:\n%s", want, got)
	}
}
