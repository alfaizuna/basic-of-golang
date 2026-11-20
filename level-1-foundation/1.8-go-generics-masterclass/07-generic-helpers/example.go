//go:build ignore

package main

import "fmt"

// Map mentransformasi setiap elemen slice []T menjadi []U
func Map[T, U any](items []T, f func(T) U) []U {
	result := make([]U, len(items))
	for i, v := range items {
		result[i] = f(v)
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	squared := Map(nums, func(n int) int {
		return n * n
	})
	fmt.Println("Squared:", squared)
}
