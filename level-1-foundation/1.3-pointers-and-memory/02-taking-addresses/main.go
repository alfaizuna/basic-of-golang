package main

import "fmt"

func main() {
	score := 42

	// 1. Ambil alamat memori score dan simpan ke variabel ptr
	// TODO: buat variabel ptr menggunakan operator &

	ptr := &score

	// 2. Cetak nilai score lewat pointer (dereference)
	// TODO: fmt.Println("Nilai awal lewat pointer:", *ptr)
	fmt.Println("Nilai awal lewat pointer:", *ptr)

	// 3. Ubah nilai score menjadi 99 lewat pointer ptr
	// TODO: *ptr = 99
	*ptr = 99

	// 4. Cetak nilai score setelah diubah
	// TODO: fmt.Println("Nilai score setelah diubah:", score)
	//fmt.Println("Nilai score setelah diubah:", *ptr)
	fmt.Println("Nilai score setelah diubah:", score)
	_ = score
	_ = fmt.Println
}
