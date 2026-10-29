//go:build ignore

package main

import "fmt"

func main() {
	// 1. Slice literal: []T tanpa ukuran tetap di dalam kurung siku
	fruits := []string{"Apel", "Jeruk", "Mangga"}

	// 2. Mengakses panjang dan indeks
	fmt.Printf("Panjang: %d, Elemen pertama: %s\n", len(fruits), fruits[0])

	// 3. Iterasi slice dengan for range
	for i, fruit := range fruits {
		fmt.Printf("[%d] %s\n", i, fruit)
	}
}
