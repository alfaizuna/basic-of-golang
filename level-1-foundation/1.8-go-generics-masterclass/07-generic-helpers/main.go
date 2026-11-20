package main

import "fmt"

// Filter menyaring slice berdasarkan fungsi predicate.
func Filter[T any](items []T, predicate func(T) bool) []T {
	// TODO: Lakukan iterasi pada items, tambahkan elemen yang memenuhi predicate ke slice baru
	return nil
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	evens := Filter(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Filtered evens: %v\n", evens)
}
