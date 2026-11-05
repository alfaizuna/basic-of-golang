//go:build ignore

package main

import "fmt"

func main() {
	// 1. Deklarasi pointer tanpa nilai awal -> zero value adalah nil
	var p *int
	fmt.Println("Zero value pointer:", p) // <nil>

	// 2. Mengecek apakah pointer bernilai nil sebelum dereference
	if p == nil {
		fmt.Println("p masih nil, jangan di-dereference!")
	}

	// 3. Menghubungkan pointer ke alamat variabel yang valid
	num := 100
	p = &num
	if p != nil {
		fmt.Println("Nilai yang ditunjuk p:", *p) // 100
	}

	// Catatan: Go sengaja TIDAK mengizinkan pointer arithmetic (seperti p++ atau p + 1)
	// demi keamanan memori!
}
