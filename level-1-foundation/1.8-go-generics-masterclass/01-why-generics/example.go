//go:build ignore

package main

import "fmt"

// Tanpa generics, fungsi harus dibuat terpisah untuk setiap tipe:
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Min int:", MinInt(10, 20))
	fmt.Println("Min float:", MinFloat(3.14, 2.71))
}
