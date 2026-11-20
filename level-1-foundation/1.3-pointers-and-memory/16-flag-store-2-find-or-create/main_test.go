package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	s := NewStore()
	flag, found := s.Get("feature-x")
	if found || flag != nil {
		t.Fatalf("expected feature-x to not be found in empty store")
	}

	created := s.FindOrCreate("feature-x")
	flag, found = s.Get("feature-x")
	if !found || flag == nil {
		t.Fatalf("expected feature-x to be found after creation")
	}
	if flag != created {
		t.Errorf("expected Get to return the exact pointer returned by FindOrCreate")
	}
}

func TestFindOrCreate(t *testing.T) {
	s := NewStore()
	f1 := s.FindOrCreate("new-checkout")
	if f1 == nil {
		t.Fatal("expected non-nil pointer from FindOrCreate")
	}
	if f1.Name != "new-checkout" || f1.Enabled != false || f1.Rollout != 0 {
		t.Errorf("unexpected default values for flag: %+v", f1)
	}

	if s.Count() != 1 {
		t.Errorf("expected store count 1, got %d", s.Count())
	}

	// Panggilan kedua harus mengembalikan alamat pointer yang persis sama
	f2 := s.FindOrCreate("new-checkout")
	if f1 != f2 {
		t.Errorf("expected f1 and f2 to point to identical memory address")
	}
	if s.Count() != 1 {
		t.Errorf("expected store count to remain 1, got %d", s.Count())
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
	wantCount := "Flag count: 1"
	wantDetail := "Flag 1: dark-mode (enabled: false, rollout: 0%)"
	wantSame := "Same memory address: true"

	if !strings.Contains(got, wantCount) {
		t.Errorf("output missing %q, got:\n%s", wantCount, got)
	}
	if !strings.Contains(got, wantDetail) {
		t.Errorf("output missing %q, got:\n%s", wantDetail, got)
	}
	if !strings.Contains(got, wantSame) {
		t.Errorf("output missing %q, got:\n%s", wantSame, got)
	}
}
