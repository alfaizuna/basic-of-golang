package main

import (
	"sync"
	"testing"
)

func TestSafeTrackerConcurrency(t *testing.T) {
	tracker := NewSafeTracker()
	var wg sync.WaitGroup
	const iterations = 100

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tracker.Visit("https://shared.com")
		}()
	}

	wg.Wait()
	if tracker.Count() != 1 {
		t.Errorf("expected count 1, got %d", tracker.Count())
	}
}
