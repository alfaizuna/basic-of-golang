//go:build ignore

package main

import "fmt"

func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println("Sum 3 angka:", Sum(1, 2, 3))

	// Unpacking slice yang sudah ada dengan '...'
	slice := []int{10, 20, 30}
	fmt.Println("Sum slice:", Sum(slice...))
}
