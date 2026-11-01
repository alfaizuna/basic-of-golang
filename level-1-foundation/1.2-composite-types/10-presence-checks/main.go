package main

import "fmt"

func checkStock(inventory map[string]int, item string) {
	// TODO: gunakan stock, exists := inventory[item]
	stock, exists := inventory[item]
	// jika exists: fmt.Printf("%s tersedia: %d unit\n", item, stock)
	// jika tidak:  fmt.Printf("%s tidak ditemukan di katalog\n", item)
	if exists {
		fmt.Printf("%s tersedia: %d unit\n", item, stock)
	} else {
		fmt.Printf("%s tidak ditemukan di katalog\n", item)
	}
}

func main() {
	// TODO:
	// 1. Buat map katalog := map[string]int{"Apel": 5, "Mangga": 0}
	katalog := map[string]int{"Apel": 5, "Mangga": 0}
	// 2. Panggil checkStock(katalog, "Mangga")
	checkStock(katalog, "Mangga")
	// 3. Panggil checkStock(katalog, "Jeruk")
	checkStock(katalog, "Jeruk")

	_ = fmt.Println
}
