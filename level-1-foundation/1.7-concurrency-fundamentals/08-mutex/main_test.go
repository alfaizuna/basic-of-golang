package main

import (
	"sync"
	"testing"
)

func TestSafeCounter(t *testing.T) {
	c := &SafeCounter{}
	var wg sync.WaitGroup
	const total = 500

	for i := 0; i < total; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	if got := c.Value(); got != total {
		t.Errorf("expected value %d, got %d", total, got)
	}
}
