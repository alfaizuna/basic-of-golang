package main

import "fmt"

// ValidateEmail memvalidasi struktur email sederhana.
func ValidateEmail(email string) bool {
	// TODO: Validasi panjang >= 5, memiliki '@' dan '.'
	return false
}

func main() {
	valid := ValidateEmail("dev@golang.org")
	fmt.Printf("Email valid: %t\n", valid)
}
