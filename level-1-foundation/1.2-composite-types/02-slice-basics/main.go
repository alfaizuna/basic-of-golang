package main

import "fmt"

// sumSlice menjumlahkan seluruh angka di dalam slice numbers.
func sumSlice(numbers []int) int {
	total := 0
	// TODO: gunakan for _, num := range numbers untuk menghitung total penjumlahan
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	// 1. Slice prices dengan 4 elemen
	prices := []int{15, 25, 35, 45}

	// 2. Cetak jumlah elemen slice menggunakan len(prices)
	fmt.Println("Count:", len(prices))

	// 3. Panggil sumSlice(prices) dan cetak hasilnya
	fmt.Println("Total:", sumSlice(prices))
}
