//go:build ignore

package main

import "fmt"

func main() {
	age := 25

	// 1. Mengambil alamat memori dengan operator & (address-of)
	ptr := &age
	fmt.Println("Alamat memori age:", ptr) // Contoh: 0x14000122010

	// 2. Membaca nilai dari alamat tersebut dengan operator * (dereference)
	fmt.Println("Nilai age lewat pointer:", *ptr) // 25

	// 3. Mengubah nilai langsung di alamat memori tersebut
	*ptr = 30
	fmt.Println("Nilai age setelah diubah:", age) // 30
}
