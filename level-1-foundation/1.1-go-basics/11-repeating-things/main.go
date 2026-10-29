package main

import "fmt"

// 1. Buat fungsi sumEven(limit int) int
// Gunakan for klasik dari 1 sampai limit (i <= limit):
// - Jika i ganjil (i % 2 != 0), lewati dengan continue
// - Jumlahkan angka genap ke variabel total
// - Kembalikan total
func sumEven(limit int) int {
	total := 0
	for i := 1; i <= limit; i++ {
		if i % 2 != 0 {
			continue
		}
		total += i
	}
	return total
}

// 2. Buat fungsi countUntil(target int) int
// Gunakan infinite loop for { ... }:
// - Tambah count dengan 3 dan tambah steps dengan 1
// - Jika count >= target, berhenti dengan break
// - Kembalikan steps
func countUntil(target int) int {
	count := 0
	steps := 0
	for {
		count += 3
		steps ++
		if count >= target {
			break
		}
	}
	return steps
}

func main() {
	// 3. Panggil sumEven(10) dan cetak hasilnya
	fmt.Println(sumEven(10))

	// 4. Panggil countUntil(10) dan cetak hasilnya
	fmt.Println(countUntil(10))
}
