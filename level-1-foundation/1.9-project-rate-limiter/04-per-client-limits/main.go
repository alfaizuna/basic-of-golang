package main

import (
	"fmt"
	"time"
)

type ClientState struct {
	count       int
	windowStart time.Time
}

type MultiClientLimiter struct {
	limit   int
	window  time.Duration
	clients map[string]*ClientState
	now     func() time.Time
}

func NewMultiClientLimiter(limit int, window time.Duration, now func() time.Time) *MultiClientLimiter {
	if now == nil {
		now = time.Now
	}
	return &MultiClientLimiter{
		limit:   limit,
		window:  window,
		clients: make(map[string]*ClientState),
		now:     now,
	}
}

// TODO: Implementasikan Allow(clientKey string) bool untuk mengisolasi kuota tiap klien
func (l *MultiClientLimiter) Allow(clientKey string) bool {
	return false
}

func main() {
	l := NewMultiClientLimiter(1, time.Minute, nil)
	fmt.Printf("User A 1: %t\n", l.Allow("userA"))
	fmt.Printf("User A 2: %t\n", l.Allow("userA"))
	fmt.Printf("User B 1: %t\n", l.Allow("userB"))
}
