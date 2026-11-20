package main

import "testing"

func TestResolveNamespace(t *testing.T) {
	if got := ResolveNamespace("myjson", "encoding/json"); got != "using alias myjson for encoding/json" {
		t.Errorf("got %q", got)
	}
	if got := ResolveNamespace("", "fmt"); got != "using default package fmt" {
		t.Errorf("got %q", got)
	}
}
