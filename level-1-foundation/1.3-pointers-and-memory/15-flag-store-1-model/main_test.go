package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestFlagStruct(t *testing.T) {
	fl := Flag{
		Name:    "beta-feature",
		Enabled: true,
		Rollout: 50,
	}
	if fl.Name != "beta-feature" || !fl.Enabled || fl.Rollout != 50 {
		t.Errorf("Flag struct field values mismatch: %+v", fl)
	}
}

func TestNewStore(t *testing.T) {
	s := NewStore()
	if s == nil {
		t.Fatal("expected NewStore to return non-nil pointer")
	}
	if s.flags == nil {
		t.Fatal("expected s.flags map to be initialized")
	}
	if s.Count() != 0 {
		t.Errorf("expected initial count to be 0, got %d", s.Count())
	}
}

func TestStoreCountWithEntries(t *testing.T) {
	s := NewStore()
	s.flags["f1"] = &Flag{Name: "f1", Enabled: true, Rollout: 100}
	s.flags["f2"] = &Flag{Name: "f2", Enabled: false, Rollout: 0}
	if s.Count() != 2 {
		t.Errorf("expected count 2, got %d", s.Count())
	}

	var nilStore *Store
	if nilStore.Count() != 0 {
		t.Errorf("expected nil store count to be 0, got %d", nilStore.Count())
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
	want := "Store created. Initial flag count: 0"
	if !strings.Contains(got, want) {
		t.Errorf("output does not contain %q, got:\n%s", want, got)
	}
}
