//go:build ignore

package main

import "fmt"

func main() {
	// 1. Array literal dengan ukuran tetap
	var a [3]int = [3]int{1, 2, 3}

	// 2. Menggunakan [...] agar compiler menghitung panjangnya
	b := [...]string{"Go", "Rust", "Python"}
	fmt.Printf("Tipe a: %T | Panjang: %d\n", a, len(a))
	fmt.Printf("Tipe b: %T | Isi: %v\n", b, b)

	// 3. Array adalah value type (disalin saat di-assign)
	c := a
	c[0] = 100
	fmt.Println("a (asli)   :", a) // [1 2 3] (tidak berubah!)
	fmt.Println("c (salinan):", c) // [100 2 3]
}
