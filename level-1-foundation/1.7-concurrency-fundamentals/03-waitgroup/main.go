package main

import (
	"fmt"
	"sync"
)

// ProcessBatch menjalankan n worker konkuren yang masing-masing menambahkan 10 ke total.
func ProcessBatch(n int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var total int

	_ = wg
	_ = mu

	// TODO: Jalankan n goroutine dengan wg.Add(1), defer wg.Done(), mu.Lock(), total += 10, mu.Unlock()
	// TODO: Tunggu semua goroutine dengan wg.Wait()

	return total
}

func main() {
	result := ProcessBatch(5)
	fmt.Printf("Total batch: %d\n", result)
}
