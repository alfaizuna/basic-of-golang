package main

import "fmt"

// ToSlice mengonversi argumen variadik bertipe generik menjadi slice.
func ToSlice[T any](items ...T) []T {
	// TODO: Kembalikan items sebagai slice
	return nil
}

func main() {
	nums := ToSlice(10, 20, 30)
	fmt.Printf("Slice: %v\n", nums)
}
