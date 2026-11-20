package main

import "testing"

func TestCalculateStats(t *testing.T) {
	min, max, total := CalculateStats([]int{10, -5, 20, 3})
	if min != -5 || max != 20 || total != 28 {
		t.Errorf("got (%d, %d, %d), want (-5, 20, 28)", min, max, total)
	}

	zMin, zMax, zTotal := CalculateStats(nil)
	if zMin != 0 || zMax != 0 || zTotal != 0 {
		t.Errorf("empty slice should return zeroes")
	}
}
