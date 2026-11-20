//go:build ignore

package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup

	// CONTOH DATA RACE: dua goroutine membaca dan menulis ke counter tanpa proteksi
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // Tidak aman untuk konkurensi!
		}()
	}

	wg.Wait()
	fmt.Println("Counter akhir (mungkin salah):", counter)
}
