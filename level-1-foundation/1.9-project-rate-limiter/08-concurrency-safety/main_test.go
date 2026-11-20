package main

import (
	"sync"
	"testing"
)

func TestSafeRateLimiterConcurrency(t *testing.T) {
	limiter := NewSafeRateLimiter(50)
	var wg sync.WaitGroup
	const goroutines = 200

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			limiter.Allow("shared-client")
		}()
	}

	wg.Wait()
	count := limiter.GetCount("shared-client")
	if count != 50 {
		t.Errorf("expected count 50, got %d", count)
	}
}
