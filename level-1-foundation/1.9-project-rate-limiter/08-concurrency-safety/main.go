package main

import (
	"fmt"
	"sync"
)

type SafeRateLimiter struct {
	mu     sync.Mutex
	limit  int
	counts map[string]int
}

func NewSafeRateLimiter(limit int) *SafeRateLimiter {
	return &SafeRateLimiter{
		limit:  limit,
		counts: make(map[string]int),
	}
}

// Allow mengevaluasi dan menambahkan request untuk key secara thread-safe.
func (s *SafeRateLimiter) Allow(key string) bool {
	// TODO: Kunci mutex dan evaluasi s.counts[key] < s.limit
	return false
}

// GetCount membaca hitungan key secara thread-safe.
func (s *SafeRateLimiter) GetCount(key string) int {
	// TODO: Kunci mutex dan kembalikan s.counts[key]
	return 0
}

func main() {
	limiter := NewSafeRateLimiter(10)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			limiter.Allow("client-1")
		}()
	}

	wg.Wait()
	fmt.Printf("Total allowed for client-1: %d\n", limiter.GetCount("client-1"))
}
