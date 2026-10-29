package main

import "fmt"

func main() {
	totalItems := 7
	bags := 2
	price := 12.5
	qty := 4

	// 1. Hitung itemsPerBag (pembagian integer totalItems / bags)
	itemsPerBag := totalItems / bags

	// 2. Hitung remainder (sisa bagi totalItems % bags)
	remainder := totalItems % bags

	// 3. Hitung avgItems (konversi totalItems dan bags ke float64, lalu bagi)
	avgItems := float64(totalItems) / float64(bags)
	// 4. Hitung totalPrice (price dikali qty yang sudah dikonversi ke float64)
	totalPrice := price * float64(qty)

	// 5. Cetak itemsPerBag, remainder, avgItems, dan totalPrice di baris terpisah
	fmt.Println(itemsPerBag)
	fmt.Println(remainder)
	fmt.Println(avgItems)
	fmt.Println(totalPrice)
}
