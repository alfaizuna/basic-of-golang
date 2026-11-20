package main

import (
	"fmt"
	"time"
)

type Limiter struct {
	Limit int
	Now   func() time.Time
}

// TODO: Buat NewLimiter(limit int, now func() time.Time) *Limiter
func NewLimiter(limit int, now func() time.Time) *Limiter {
	return nil
}

// TODO: Buat method Allow() bool pada *Limiter
func (l *Limiter) Allow() bool {
	return false
}

func main() {
	l := NewLimiter(5, nil)
	fmt.Printf("Limiter ready: %t\n", l.Allow())
}
