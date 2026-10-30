package main

import "fmt"

func removeAt(items []string, index int) []string {
	// TODO: kembalikan slice baru setelah elemen pada index dihapus
	return append(items[:index], items[index+1:]...)
}

func main() {
	// TODO:
	// 1. Buat slice languages := []string{"Go", "Java", "Python", "Ruby"}
	languages := []string{"Go", "Java", "Python", "Ruby"}
	// 2. Hapus indeks 1 menggunakan removeAt dan tampung kembali ke languages
	languages = removeAt(languages, 1)
	// 3. Cetak: fmt.Println("Languages:", languages)
	fmt.Println("Languages:", languages)
	// 4. Buat nil slice var empty []string
	var empty []string
	// 5. Cetak: fmt.Println("Is Nil:", empty == nil)
	fmt.Println("Is Nil:", empty == nil)

	_ = fmt.Println
}
