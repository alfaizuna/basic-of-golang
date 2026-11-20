package main

import "fmt"

// TODO: Definisikan interface RateLimiter
type RateLimiter interface {
}

type SimulatedLimiter struct {
	maxRequests int
	current     int
}

// TODO: Implementasikan method Allow(key string) bool pada *SimulatedLimiter

// TODO: Implementasikan ExecuteRequests(rl RateLimiter, key string, count int) int
func ExecuteRequests(rl RateLimiter, key string, count int) int {
	return 0
}

func main() {
	sim := &SimulatedLimiter{maxRequests: 5}
	permitted := ExecuteRequests(sim, "clientA", 10)
	fmt.Printf("Total permitted: %d of 10\n", permitted)
}
