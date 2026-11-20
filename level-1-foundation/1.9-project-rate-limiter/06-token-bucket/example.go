//go:build ignore

package main

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	capacity   float64
	tokens     float64
	refillRate float64 // token per detik
	lastRefill time.Time
}

func main() {
	fmt.Println("Token bucket algorithm")
}
