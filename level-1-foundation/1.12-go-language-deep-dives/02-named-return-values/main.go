package main

import "fmt"

// CalculateStats menghitung min, max, dan total dari sekumpulan angka.
func CalculateStats(numbers []int) (min int, max int, total int) {
	// TODO: Hitung min, max, dan total
	return 0, 0, 0
}

func main() {
	numbers := []int{5, 2, 9}
	min, max, total := CalculateStats(numbers)
	fmt.Printf("Min: %d, Max: %d, Total: %d\n", min, max, total)
}
