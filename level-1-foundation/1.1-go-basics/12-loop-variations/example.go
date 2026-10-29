//go:build ignore

package main

import "fmt"

func main() {
	// 1. for range pada string
	// Menghasilkan indeks byte dan karakter (rune)
	for _, char := range "Go" {
		fmt.Printf("Karakter: %c\n", char)
	}

	// 2. Loop bersarang (nested loop) dengan Label
	// Label memungkinkan break keluar langsung dari loop terluar
Search:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j == 4 {
				fmt.Printf("Ditemukan pertama kali pada i=%d, j=%d\n", i, j)
				break Search // langsung keluar dari loop luar (Search)
			}
		}
	}
}
