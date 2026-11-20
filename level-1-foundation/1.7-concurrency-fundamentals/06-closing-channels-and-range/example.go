//go:build ignore

package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // Menutup channel

	// Loop for-range akan berhenti otomatis saat channel ditutup dan buffer habis
	for val := range ch {
		fmt.Println("Menerima:", val)
	}

	// Pemeriksaan status channel dengan comma-ok idiom
	val, ok := <-ch
	fmt.Printf("Nilai: %d, Terbuka: %t\n", val, ok)
}
