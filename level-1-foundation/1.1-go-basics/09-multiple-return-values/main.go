package main

import "fmt"

// 1. Buat fungsi calculate(a, b int) (int, int) di sini
// yang mengembalikan (a + b) dan (a - b)
func calculate(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	// 2. Panggil calculate(20, 5) dan simpan hasilnya ke variabel sum dan diff
	// Cetak hasilnya dengan fmt.Println(sum, diff)
	sum, diff := calculate(20, 5)
	fmt.Println(sum, diff)

	// 3. Panggil calculate(50, 10), simpan hasil pertama ke variabel total
	// dan abaikan hasil kedua dengan blank identifier (_)
	// Cetak hasilnya dengan fmt.Println(total)
	total, _ := calculate(50, 10)
	fmt.Println(total)

}
