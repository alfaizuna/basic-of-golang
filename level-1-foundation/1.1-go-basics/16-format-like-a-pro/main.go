package main

import "fmt"

// 1. formatProduct memformat detail produk.
// Format yang diharapkan: "Product: %q | Price: $%.2f | Stock: %d"
func formatProduct(name string, price float64, stock int) string {
	// TODO: gunakan fmt.Sprintf dengan verb %q, %.2f, dan %d
	return fmt.Sprintf("Product: %q | Price: $%.2f | Stock: %d", name, price, stock)
}

// 2. describeValue menampilkan tipe data dan nilai yang di-padding dengan 4 digit.
// Format yang diharapkan: "Type: %T | Padded: %04d"
func describeValue(val int) string {
	// TODO: gunakan fmt.Sprintf dengan verb %T dan %04d (berikan val dua kali)
	return fmt.Sprintf("Type: %T | Padded: %04d", val, val)
}

func main() {
	// Contoh pemanggilan 1: formatProduct
	p := formatProduct("Mechanical Keyboard", 89.9, 15)
	fmt.Println(p)

	// TODO: Panggil describeValue(42) dan cetak hasilnya
	dv := describeValue(42)
	fmt.Println(dv)
}
