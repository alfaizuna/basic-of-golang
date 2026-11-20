package main

import (
	"strings"
	"testing"
)

func TestFormatModuleHeader(t *testing.T) {
	got := FormatModuleHeader("example.com/app", "1.22.0")
	expected := "module example.com/app\n\ngo 1.22.0"
	if strings.TrimSpace(got) != strings.TrimSpace(expected) {
		t.Errorf("got %q, want %q", got, expected)
	}
}
