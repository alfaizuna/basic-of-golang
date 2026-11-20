package main

import "testing"

func TestDescribePackage(t *testing.T) {
	if got := DescribePackage("main", true); got != "main (executable)" {
		t.Errorf("got %q, want 'main (executable)'", got)
	}
	if got := DescribePackage("auth", false); got != "auth (library)" {
		t.Errorf("got %q, want 'auth (library)'", got)
	}
}
