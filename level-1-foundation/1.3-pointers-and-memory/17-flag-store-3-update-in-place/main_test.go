package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestFlagEnableDisable(t *testing.T) {
	fl := &Flag{Name: "feature", Enabled: false, Rollout: 0}
	fl.Enable()
	if !fl.Enabled {
		t.Errorf("expected Enabled to be true after Enable()")
	}

	fl.Disable()
	if fl.Enabled {
		t.Errorf("expected Enabled to be false after Disable()")
	}

	// Nil safety check
	var nilFlag *Flag
	nilFlag.Enable()
	nilFlag.Disable()
}

func TestFlagSetRollout(t *testing.T) {
	fl := &Flag{Name: "feature", Enabled: true, Rollout: 10}

	err := fl.SetRollout(50)
	if err != nil {
		t.Errorf("unexpected error for valid rollout: %v", err)
	}
	if fl.Rollout != 50 {
		t.Errorf("expected Rollout 50, got %d", fl.Rollout)
	}

	// Invalid cases
	if err := fl.SetRollout(-1); err == nil {
		t.Errorf("expected error for negative rollout")
	}
	if err := fl.SetRollout(101); err == nil {
		t.Errorf("expected error for rollout > 100")
	}

	// Nil safety check
	var nilFlag *Flag
	if err := nilFlag.SetRollout(50); err == nil {
		t.Errorf("expected error when setting rollout on nil flag")
	}
}

func TestStoreUpdateInPlace(t *testing.T) {
	store := NewStore()
	flagRef := store.FindOrCreate("payment-v2")

	// Mutasi lewat pointer
	flagRef.Enable()
	_ = flagRef.SetRollout(80)

	// Verifikasi lewat store
	fromStore, ok := store.Get("payment-v2")
	if !ok || fromStore == nil {
		t.Fatal("expected flag to exist in store")
	}
	if !fromStore.Enabled || fromStore.Rollout != 80 {
		t.Errorf("store did not reflect in-place pointer mutation: %+v", fromStore)
	}
	if fromStore != flagRef {
		t.Errorf("expected fromStore and flagRef to share identical memory address")
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
	want1 := "Initial: beta-ui (enabled: false, rollout: 0%)"
	want2 := "Retrieved from store: beta-ui (enabled: true, rollout: 75%)"
	want3 := "Shared memory confirmed: true"

	if !strings.Contains(got, want1) {
		t.Errorf("output does not contain %q, got:\n%s", want1, got)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("output does not contain %q, got:\n%s", want2, got)
	}
	if !strings.Contains(got, want3) {
		t.Errorf("output does not contain %q, got:\n%s", want3, got)
	}
}
