package main

import "testing"

func TestStack(t *testing.T) {
	var s Stack[string]
	if _, ok := s.Pop(); ok {
		t.Fatal("expected empty pop to return false")
	}

	s.Push("apple")
	s.Push("banana")

	if s.Len() != 2 {
		t.Errorf("expected len 2, got %d", s.Len())
	}

	val, ok := s.Pop()
	if !ok || val != "banana" {
		t.Errorf("expected banana, got %v", val)
	}
}
