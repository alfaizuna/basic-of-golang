package main

import "testing"

func TestJoinStrings(t *testing.T) {
	if got := JoinStrings("/", "usr", "local", "bin"); got != "usr/local/bin" {
		t.Errorf("got %q, want 'usr/local/bin'", got)
	}

	items := []string{"a", "b", "c"}
	if got := JoinStrings(",", items...); got != "a,b,c" {
		t.Errorf("unpacking failed, got %q", got)
	}

	if got := JoinStrings("-"); got != "" {
		t.Errorf("empty variadic should return empty, got %q", got)
	}
}
