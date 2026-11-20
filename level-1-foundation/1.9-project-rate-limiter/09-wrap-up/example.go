//go:build ignore

package main

import "fmt"

type RateLimiter interface {
	Allow(key string) bool
}

type AlwaysAllowLimiter struct{}

func (a *AlwaysAllowLimiter) Allow(key string) bool {
	return true
}

func main() {
	var limiter RateLimiter = &AlwaysAllowLimiter{}
	fmt.Println("Allowed:", limiter.Allow("test"))
}
