package main

import "fmt"

// RequirePositive memicu panic jika nilai kurang dari atau sama dengan nol.
func RequirePositive(val int) {
	// TODO: Panggil panic jika val <= 0
}

func main() {
	RequirePositive(42)
	fmt.Println("Valid value: 42")
}
