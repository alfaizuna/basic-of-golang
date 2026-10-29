package main

import "fmt"

func main() {
	// Buat grade calculator dengan if initializer untuk score := 85:
	// - score >= 90: cetak "Grade: A"
	// - score >= 80: cetak "Grade: B"
	// - score >= 70: cetak "Grade: C"
	// - lainnya: cetak "Grade: F"

	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}
}