//go:build ignore

package main

import "fmt"

func main() {
	// 1. append pada nil slice langsung bekerja secara aman
	var numbers []int // nil slice (len: 0, cap: 0)
	fmt.Printf("len: %d, cap: %d, isi: %v\n", len(numbers), cap(numbers), numbers)
	
	numbers = append(numbers, 10)
	fmt.Printf("len: %d, cap: %d, isi: %v\n", len(numbers), cap(numbers), numbers)

	// 2. Pertumbuhan kapasitas otomatis saat slice penuh
	numbers = append(numbers, 20)
	numbers = append(numbers, 30) // kapasitas bertambah otomatis di balik layar
	fmt.Printf("len: %d, cap: %d, isi: %v\n", len(numbers), cap(numbers), numbers)

	// 3. Menambahkan beberapa elemen sekaligus
	numbers = append(numbers, 40, 50)
	fmt.Printf("len: %d, cap: %d, isi: %v\n", len(numbers), cap(numbers), numbers)
}
