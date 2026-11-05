package main

import "fmt"

func main() {
	// 1. Deklarasikan variabel pointer ptr dengan tipe *string tanpa inisialisasi
	// TODO: var ptr *string
	var ptr *string

	// 2. Cek jika ptr bernilai nil, lalu cetak "ptr adalah nil"
	// TODO: if ptr == nil { ... }
	if ptr == nil {
		fmt.Println("ptr adalah nil")
	}

	message := "Belajar Go menyenangkan!"

	// 3. Pasangkan alamat message ke ptr
	// TODO: ptr = &message
	ptr = &message

	// 4. Cek jika ptr tidak nil, lalu cetak nilai isinya lewat dereference:
	// "ptr tidak nil, isinya:", *ptr
	// TODO: if ptr != nil { ... }
	if ptr != nil {
		fmt.Println("ptr tidak nil, isinya:", *ptr)
	}

	_ = message
	_ = fmt.Println
}
