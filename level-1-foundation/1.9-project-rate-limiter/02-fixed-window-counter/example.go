//go:build ignore

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

func main() {
	fmt.Println("Fixed window limiter initialized")
}
