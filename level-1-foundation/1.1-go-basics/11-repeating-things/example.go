//go:build ignore

package main

import "fmt"

func main() {
	// Bentuk 1: Klasik (init; condition; post)
	for i := 1; i <= 3; i++ {
		fmt.Println("Klasik:", i)
	}

	// Bentuk 2: Mirip "while" (hanya kondisi)
	count := 1
	for count <= 3 {
		fmt.Println("While-style:", count)
		count++
	}

	// Bentuk 3: Infinite loop dengan break dan continue
	num := 0
	for {
		num++
		if num == 2 {
			continue // lewati angka 2
		}
		if num > 4 {
			break // berhenti saat num > 4
		}
		fmt.Println("Infinite loop:", num)
	}
}
