package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestResolveChain(t *testing.T) {
	z := NewZone("example.com")
	target := NewRecord("origin.example.com", TypeA, "1.2.3.4", 300)
	z.AddRecord(target)
	z.AddRecord(NewRecord("alias1.example.com", TypeCNAME, "alias2.example.com", 300))
	z.AddRecord(NewRecord("alias2.example.com", TypeCNAME, "origin.example.com", 300))

	// Resolve alias1 -> alias2 -> origin
	res, err := z.Resolve("alias1.example.com")
	if err != nil {
		t.Fatalf("unexpected error resolving chain: %v", err)
	}
	if res != target {
		t.Errorf("expected to resolve to target record, got %v", res)
	}

	// Resolve direct A record
	resDirect, err := z.Resolve("origin.example.com")
	if err != nil || resDirect != target {
		t.Errorf("expected direct resolution to succeed")
	}

	// Resolve missing domain
	_, errMissing := z.Resolve("missing.example.com")
	if errMissing == nil {
		t.Errorf("expected error for non-existent domain")
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
	want := "Resolved app.example.com -> 93.184.216.34"

	if !strings.Contains(got, want) {
		t.Errorf("output does not contain %q, got:\n%s", want, got)
	}
}
