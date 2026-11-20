package main

import "testing"

func TestSimulateWorkload(t *testing.T) {
	cases := []struct {
		workers  int
		expected int
	}{
		{0, 0},
		{1, 10},
		{4, 40},
		{10, 100},
	}
	for _, tc := range cases {
		got := SimulateWorkload(tc.workers)
		if got != tc.expected {
			t.Errorf("SimulateWorkload(%d) = %d, want %d", tc.workers, got, tc.expected)
		}
	}
}
