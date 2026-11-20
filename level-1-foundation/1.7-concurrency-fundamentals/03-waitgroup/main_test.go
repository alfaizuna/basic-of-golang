package main

import "testing"

func TestProcessBatch(t *testing.T) {
	cases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 10},
		{5, 50},
		{100, 1000},
	}
	for _, tc := range cases {
		got := ProcessBatch(tc.n)
		if got != tc.expected {
			t.Errorf("ProcessBatch(%d) = %d, want %d", tc.n, got, tc.expected)
		}
	}
}
