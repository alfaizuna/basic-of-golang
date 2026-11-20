package main

import "testing"

func TestAccumulateConcurrent(t *testing.T) {
	numbers := make([]int, 100)
	var expected int
	for i := 0; i < 100; i++ {
		numbers[i] = i + 1
		expected += (i + 1)
	}

	got := AccumulateConcurrent(numbers)
	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
