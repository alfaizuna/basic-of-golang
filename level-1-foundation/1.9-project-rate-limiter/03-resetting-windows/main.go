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

// TODO: Perbarui method Allow() agar mereset windowStart dan count saat cur.Sub(l.windowStart) >= l.window
func (l *FixedWindowLimiter) Allow() bool {
	return false
}

func main() {
	fakeTime := time.Now()
	nowFn := func() time.Time { return fakeTime }

	l := NewFixedWindowLimiter(1, 10*time.Second, nowFn)
	fmt.Printf("Req 1: %t\n", l.Allow())
	fmt.Printf("Req 2: %t\n", l.Allow())

	fakeTime = fakeTime.Add(11 * time.Second)
	fmt.Printf("After window reset: %t\n", l.Allow())
}
