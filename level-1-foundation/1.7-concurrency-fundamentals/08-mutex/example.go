//go:build ignore

package main

import (
	"fmt"
	"sync"
)

type SafeStore struct {
	mu   sync.Mutex
	data map[string]int
}

func (s *SafeStore) Set(key string, val int) {
	s.mu.Lock()
	defer s.mu.Unlock() // Dilepaskan saat fungsi Set selesai
	s.data[key] = val
}

func (s *SafeStore) Get(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data[key]
}

func main() {
	store := SafeStore{data: make(map[string]int)}
	store.Set("gopher", 100)
	fmt.Println("Stored:", store.Get("gopher"))
}
