//go:build ignore

package main

import "fmt"

// Fungsi yang mengembalikan 2 nilai: hasil bagi dan sisa bagi
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	// Menampung kedua nilai kembalian
	q, r := divide(10, 3)
	fmt.Printf("Hasil: %d, Sisa: %d\n", q, r)

	// Mengabaikan sisa bagi menggunakan blank identifier (_)
	onlyQuotient, _ := divide(10, 3)
	fmt.Println("Hanya hasil:", onlyQuotient)
}
