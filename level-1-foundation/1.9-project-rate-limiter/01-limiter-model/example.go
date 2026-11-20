//go:build ignore

package main

import (
	"fmt"
	"time"
)

type Limiter struct {
	limit int
	now   func() time.Time // Clock injection
}

func (l *Limiter) Allow() bool {
	return true
}

func main() {
	limiter := &Limiter{
		limit: 10,
		now:   time.Now,
	}
	fmt.Println("Allowed:", limiter.Allow())
}
