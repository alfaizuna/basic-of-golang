package main

import "testing"

func TestPair(t *testing.T) {
	p := NewPair("User", 100)
	if p.First != "User" || p.Second != 100 {
		t.Errorf("unexpected pair contents: %v", p)
	}
}
