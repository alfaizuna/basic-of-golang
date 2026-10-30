//go:build ignore

package main

import "fmt"

func main() {
	// 1. Nil slice vs Empty slice
	var nilSlice []int    // len: 0, cap: 0, nilSlice == nil -> true
	emptySlice := []int{} // len: 0, cap: 0, emptySlice == nil -> false

	fmt.Printf("nilSlice == nil: %v, len: %d\n", nilSlice == nil, len(nilSlice))
	fmt.Printf("emptySlice == nil: %v, len: %d\n", emptySlice == nil, len(emptySlice))

	// Keduanya sama-sama aman untuk append dan len
	nilSlice = append(nilSlice, 10)
	fmt.Println("nilSlice setelah append:", nilSlice)

	// 2. Resep: Menghapus elemen pada indeks i
	// Gunakan: append(s[:i], s[i+1:]...)
	letters := []string{"A", "B", "C", "D"}
	// Hapus "B" di indeks 1
	letters = append(letters[:1], letters[2:]...)
	fmt.Println("Setelah hapus indeks 1:", letters)
}
