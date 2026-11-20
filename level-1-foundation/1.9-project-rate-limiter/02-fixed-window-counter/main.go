package main

import (
	"fmt"
	"time"
)

type FixedWindowLimiter struct {
	limit       int
	window      time.Duration
	count       int
	windowStart time.Time
	now         func() time.Time
}

func NewFixedWindowLimiter(limit int, window time.Duration, now func() time.Time) *FixedWindowLimiter {
	if now == nil {
		now = time.Now
	}
	return &FixedWindowLimiter{
		limit:       limit,
		window:      window,
		windowStart: now(),
		now:         now,
	}
}

// TODO: Implementasikan method Allow() bool
func (l *FixedWindowLimiter) Allow() bool {
	return false
}

func main() {
	l := NewFixedWindowLimiter(2, time.Minute, nil)
	fmt.Printf("Allowed 1: %t\n", l.Allow())
	fmt.Printf("Allowed 2: %t\n", l.Allow())
	fmt.Printf("Allowed 3: %t\n", l.Allow())
}
