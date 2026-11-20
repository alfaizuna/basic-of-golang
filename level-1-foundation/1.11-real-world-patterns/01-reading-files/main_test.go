package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfigText(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "app.conf")
	_ = os.WriteFile(f, []byte("port=8080"), 0644)

	content, err := LoadConfigText(f)
	if err != nil || content != "port=8080" {
		t.Fatalf("unexpected result: content=%q, err=%v", content, err)
	}

	_, err = LoadConfigText(filepath.Join(dir, "missing.conf"))
	if err == nil || err.Error() != "file not found" {
		t.Errorf("expected 'file not found' error, got %v", err)
	}
}
