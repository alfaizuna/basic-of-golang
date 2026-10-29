//go:build ignore

package main

import (
	"errors"
	"fmt"
)

// Fungsi yang berpotensi gagal mengembalikan error di posisi return terakhir
func checkAge(age int) (string, error) {
	if age < 0 {
		return "", errors.New("age cannot be negative")
	}
	if age >= 18 {
		return "Adult", nil
	}
	return "Minor", nil
}

func main() {
	// Pola idiomatik Go: if err != nil
	status, err := checkAge(-5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Status:", status)
}
