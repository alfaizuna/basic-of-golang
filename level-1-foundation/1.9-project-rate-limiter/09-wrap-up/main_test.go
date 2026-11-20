package main

import "testing"

func TestRateLimiterInterface(t *testing.T) {
	sim := &SimulatedLimiter{maxRequests: 3}
	var rl RateLimiter = sim

	allowed := ExecuteRequests(rl, "test", 5)
	if allowed != 3 {
		t.Errorf("expected 3 allowed, got %d", allowed)
	}
}
