//go:build ignore

package main

import "fmt"

func main() {
	// 1. make dengan 2 argumen: make([]T, len) -> cap otomatis sama dengan len
	zeros := make([]int, 3)
	fmt.Printf("zeros  -> len: %d, cap: %d, isi: %v\n", len(zeros), cap(zeros), zeros)

	// 2. make dengan 3 argumen: make([]T, len, cap)
	buffer := make([]string, 1, 4)
	buffer[0] = "Ready"
	fmt.Printf("buffer -> len: %d, cap: %d, isi: %v\n", len(buffer), cap(buffer), buffer)
}
