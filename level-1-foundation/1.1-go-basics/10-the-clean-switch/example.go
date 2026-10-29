//go:build ignore

package main

import "fmt"

func main() {
	role := "admin"

	// 1. Switch biasa (dengan ekspresi nilai)
	switch role {
	case "admin":
		fmt.Println("Akses penuh")
	case "editor", "author": // multiple values dalam satu case
		fmt.Println("Akses edit")
	default:
		fmt.Println("Akses tamu")
	}

	score := 85

	// 2. Switch tanpa ekspresi (seperti rantai if-else yang rapi)
	switch {
	case score >= 90:
		fmt.Println("Grade A")
	case score >= 80:
		fmt.Println("Grade B")
	default:
		fmt.Println("Grade C")
	}
}
