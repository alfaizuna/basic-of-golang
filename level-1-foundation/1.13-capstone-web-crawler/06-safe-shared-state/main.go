package main

import (
	"fmt"
	"sync"
)

type SafeTracker struct {
	mu      sync.Mutex
	visited map[string]bool
}

func NewSafeTracker() *SafeTracker {
	// TODO: Inisialisasi struct SafeTracker dengan map visited
	return &SafeTracker{visited: make(map[string]bool)}
}

func (s *SafeTracker) Visit(url string) bool {
	// TODO: Implementasikan Visit secara thread-safe
	return false
}

func (s *SafeTracker) Count() int {
	// TODO: Implementasikan Count secara thread-safe
	return 0
}

func main() {
	tracker := NewSafeTracker()
	res := tracker.Visit("https://store.local")
	fmt.Printf("Visited: %t\n", res)
	fmt.Printf("Total unique: %d\n", tracker.Count())
}
