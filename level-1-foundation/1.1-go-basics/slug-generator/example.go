//go:build ignore

package main

import (
	"errors"
	"fmt"
	"strings"
)

// Contoh fungsi yang memvalidasi input dan mengembalikan error jika kosong
func validateUsername(name string) (string, error) {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return "", errors.New("username cannot be empty")
	}
	return strings.ToLower(trimmed), nil
}

func main() {
	user, err := validateUsername("  Alfaizuna  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Username valid: %q\n", user)

	_, err2 := validateUsername("   ")
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
}
