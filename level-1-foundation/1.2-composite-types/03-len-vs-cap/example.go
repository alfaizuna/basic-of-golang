//go:build ignore

package main

import "fmt"

func main() {
	// Slice awal dengan 5 elemen
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("numbers -> len: %d, cap: %d\n", len(numbers), cap(numbers)) // len: 5, cap: 5

	// Sub-slice dari indeks 2 sampai sebelum 4
	sub := numbers[2:4]
	// len = 2 (elemen 3 dan 4)
	// cap = 3 (indeks 2 s/d 4 di backing array: 3, 4, 5)
	fmt.Printf("sub     -> len: %d, cap: %d\n", len(sub), cap(sub))         // len: 2, cap: 3
}
