//go:build ignore

package main

import "fmt"

func main() {
	src := []int{1, 2, 3}

	// dst HARUS sudah memiliki length yang cukup!
	dst := make([]int, len(src))
	n := copy(dst, src)

	fmt.Printf("Jumlah elemen disalin: %d\n", n)
	fmt.Println("dst:", dst)

	// Buktikan independen: ubah dst, src tidak berubah
	dst[0] = 999
	fmt.Println("\nSetelah dst[0] = 999:")
	fmt.Println("src:", src)
	fmt.Println("dst:", dst)
}
