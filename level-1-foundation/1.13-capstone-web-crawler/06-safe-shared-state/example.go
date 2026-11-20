//go:build ignore

package main

import (
	"fmt"
	"sync"
)

type SafeTracker struct {
	mu      sync.Mutex
	visited map[string]bool
}

func (s *SafeTracker) CheckAndMark(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.visited[url] {
		return false // sudah pernah dikunjungi
	}
	s.visited[url] = true
	return true // pertama kali
}

func main() {
	tracker := SafeTracker{visited: make(map[string]bool)}
	fmt.Println("First time:", tracker.CheckAndMark("https://a.com"))
	fmt.Println("Second time:", tracker.CheckAndMark("https://a.com"))
}
