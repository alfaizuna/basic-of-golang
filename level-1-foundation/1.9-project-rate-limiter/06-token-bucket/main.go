package main

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	capacity   float64
	tokens     float64
	refillRate float64
	lastRefill time.Time
	now        func() time.Time
}

func NewTokenBucket(capacity float64, refillRate float64, now func() time.Time) *TokenBucket {
	if now == nil {
		now = time.Now
	}
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: now(),
		now:        now,
	}
}

// TODO: Implementasikan method Allow() bool sesuai algoritma Token Bucket
func (tb *TokenBucket) Allow() bool {
	return false
}

func main() {
	curr := time.Now()
	tb := NewTokenBucket(2.0, 1.0, func() time.Time { return curr })

	fmt.Printf("Token 1: %t\n", tb.Allow())
	fmt.Printf("Token 2: %t\n", tb.Allow())
	fmt.Printf("Token 3: %t\n", tb.Allow())

	curr = curr.Add(1 * time.Second)
	fmt.Printf("After refill: %t\n", tb.Allow())
}
