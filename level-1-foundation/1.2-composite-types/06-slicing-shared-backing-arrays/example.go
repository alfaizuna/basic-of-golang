//go:build ignore

package main

import "fmt"

func main() {
	original := []int{10, 20, 30, 40, 50}

	// sub mengambil indeks 1 sampai 3 (tidak termasuk 3)
	sub := original[1:3] // [20, 30]
	fmt.Println("Original sebelum:", original)
	fmt.Println("Sub:", sub)

	// Mengubah elemen di sub berdampak langsung ke original!
	sub[0] = 999
	fmt.Println("\nSetelah sub[0] = 999:")
	fmt.Println("Sub:", sub)
	fmt.Println("Original:", original)
}
