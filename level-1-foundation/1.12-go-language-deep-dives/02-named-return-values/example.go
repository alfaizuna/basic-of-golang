//go:build ignore

package main

import "fmt"

// Named return values: min dan max otomatis terdefinisi sebagai variabel lokal int
func MinMax(a, b int) (min, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return min, max // Eksplisit lebih disukai daripada naked return
}

func main() {
	low, high := MinMax(15, 8)
	fmt.Printf("Min: %d, Max: %d\n", low, high)
}
