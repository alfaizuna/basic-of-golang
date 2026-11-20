//go:build ignore

package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	mu   sync.Mutex
	data map[string]int
}

func (c *ConcurrentMap) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key]++
}

func main() {
	m := ConcurrentMap{data: make(map[string]int)}
	m.Inc("hit")
	fmt.Println("Hit count:", m.data["hit"])
}
