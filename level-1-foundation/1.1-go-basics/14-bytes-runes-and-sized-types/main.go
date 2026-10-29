package main

import "fmt"

// 1. toByte memeriksa apakah rune muat dalam 1 byte (<= 255).
// Jika r <= 255, kembalikan byte(r) dan true.
// Jika tidak, kembalikan 0 dan false.
func toByte(r rune) (byte, bool) {
	if r <= 255 {
		// TODO: kembalikan byte(r) dan true (saat ini masih 0, false)
		return byte(r), true
	}
	return 0, false
}

// 2. addInt8 menjumlahkan dua bilangan int8.
// Perhatikan efek overflow saat nilainya melewati 127.
func addInt8(a int8, b int8) int8 {
	// TODO: kembalikan hasil penjumlahan a + b
	return a + b
}

func main() {
	// Contoh pemanggilan 1: karakter ASCII 'A'
	b1, ok1 := toByte('A')
	fmt.Println(b1, ok1)

	// TODO: Panggil toByte('世') dan cetak hasilnya (b2, ok2)
	b2, ok2 := toByte('世')
	fmt.Println(b2, ok2)

	// TODO: Panggil addInt8(127, 1) dan cetak hasilnya
	fmt.Println(addInt8(127,1))
}
