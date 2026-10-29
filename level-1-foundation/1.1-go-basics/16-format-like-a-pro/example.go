//go:build ignore

package main

import "fmt"

func main() {
	nama := "Budi"
	umur := 25
	saldo := 150000.75
	isAktif := true

	// 1. Verb umum: %v (default), %T (tipe data)
	fmt.Printf("Default: %v | Tipe: %T\n", saldo, saldo)

	// 2. Verb spesifik: %s (string), %q (quoted), %d (integer), %t (boolean)
	fmt.Printf("Nama: %s (atau %q) | Umur: %d | Aktif: %t\n", nama, nama, umur, isAktif)

	// 3. Lebar dan Presisi: %.2f (2 digit desimal), %04d (padding nol)
	fmt.Printf("Saldo: Rp %.2f | ID: %04d\n", saldo, 7)

	// 4. fmt.Sprintf mengembalikan string tanpa langsung mencetaknya ke konsol
	pesan := fmt.Sprintf("Halo %s, saldo Anda: Rp %.2f", nama, saldo)
	fmt.Println("Hasil Sprintf:", pesan)
}
