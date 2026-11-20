package main

import (
	"fmt"
	"time"
)

type SlidingWindowLog struct {
	limit      int
	window     time.Duration
	timestamps []time.Time
	now        func() time.Time
}

func NewSlidingWindowLog(limit int, window time.Duration, now func() time.Time) *SlidingWindowLog {
	if now == nil {
		now = time.Now
	}
	return &SlidingWindowLog{
		limit:  limit,
		window: window,
		now:    now,
	}
}

// TODO: Implementasikan Allow() bool dengan memangkas timestamp lama dan mengecek batas limit
func (s *SlidingWindowLog) Allow() bool {
	return false
}

func main() {
	curr := time.Now()
	s := NewSlidingWindowLog(2, 5*time.Second, func() time.Time { return curr })

	fmt.Printf("Req 1: %t\n", s.Allow())
	fmt.Printf("Req 2: %t\n", s.Allow())
	fmt.Printf("Req 3: %t\n", s.Allow())

	curr = curr.Add(6 * time.Second)
	fmt.Printf("After slide: %t\n", s.Allow())
}
