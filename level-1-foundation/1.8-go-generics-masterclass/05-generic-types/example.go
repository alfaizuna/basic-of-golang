//go:build ignore

package main

import "fmt"

// Pair menampung dua nilai dengan tipe yang bisa berbeda
type Pair[T, U any] struct {
	First  T
	Second U
}

func main() {
	// Inisialisasi struct generik dengan menyebutkan tipenya secara eksplisit
	p1 := Pair[string, int]{First: "Umur", Second: 25}
	p2 := Pair[int, bool]{First: 200, Second: true}

	fmt.Printf("p1: (%s, %d)\n", p1.First, p1.Second)
	fmt.Printf("p2: (%d, %t)\n", p2.First, p2.Second)
}
