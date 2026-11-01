//go:build ignore

package main

import "fmt"

func main() {
	// 1. Membuat map dengan make
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30

	// 2. Map literal
	prices := map[string]int{
		"Kopi": 15000,
		"Teh":  10000,
	}

	// 3. Mengakses dan menghapus
	fmt.Println("Harga Kopi:", prices["Kopi"])
	delete(prices, "Teh")

	fmt.Println("Jumlah item prices:", len(prices))
	fmt.Println("Isi prices:", prices)
}
