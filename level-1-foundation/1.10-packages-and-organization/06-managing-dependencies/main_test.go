package main

import "testing"

func TestFormatSemver(t *testing.T) {
	if got := FormatSemver(1, 0, 0, ""); got != "v1.0.0" {
		t.Errorf("got %q, want 'v1.0.0'", got)
	}
	if got := FormatSemver(2, 1, 4, "rc.1"); got != "v2.1.4-rc.1" {
		t.Errorf("got %q, want 'v2.1.4-rc.1'", got)
	}
}
