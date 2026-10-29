package main

import "fmt"

// 1. Buat fungsi greet(name string) string
// yang mengembalikan "Halo, " + name
func greet(name string) string {
	return "Halo, " + name
}
// 2. Buat fungsi double(n int) int
// yang mengembalikan n * 2
func double(n int) int {
	return n * 2
}

func main() {
	// 3. Panggil greet("Go") dan cetak hasilnya
	fmt.Println(greet("Go"))
	// 4. Panggil double(21) dan cetak hasilnya
	fmt.Println(double(21))
}
