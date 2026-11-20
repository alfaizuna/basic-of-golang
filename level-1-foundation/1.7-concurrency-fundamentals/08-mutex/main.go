package main

import (
	"fmt"
	"sync"
)

// SafeCounter menyimpan counter yang aman diakses secara konkuren.
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Inc menambah value secara thread-safe.
func (c *SafeCounter) Inc() {
	// TODO: Implementasikan method Inc() dengan pointer receiver dan defer mu.Unlock()
}

// Value mengembalikan nilai value secara thread-safe.
func (c *SafeCounter) Value() int {
	// TODO: Implementasikan method Value() int dengan pointer receiver dan defer mu.Unlock()
	return 0
}

func main() {
	counter := &SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Printf("Final count: %d\n", counter.Value())
}
