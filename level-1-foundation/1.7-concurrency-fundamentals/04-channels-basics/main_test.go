package main

import "testing"

func TestComputeSquare(t *testing.T) {
	cases := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{3, 9},
		{9, 81},
		{-5, 25},
	}
	for _, tc := range cases {
		got := ComputeSquare(tc.input)
		if got != tc.expected {
			t.Errorf("ComputeSquare(%d) = %d, want %d", tc.input, got, tc.expected)
		}
	}
}
