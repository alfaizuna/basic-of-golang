package main

import "fmt"

func main() {
	original := []string{"Go", "Rust", "Python"}

	// TODO:
	// 1. Buat slice backup dengan make berukuran sama dengan len(original)
	backup := make([]string, len(original))
	// 2. Salin isi original ke backup menggunakan copy
	copy(backup, original)
	// 3. Ubah elemen terakhir backup menjadi "TypeScript"
	backup[2] = "TypeScript"
	// 4. Cetak: fmt.Println("Original:", original)
	fmt.Println("Original:", original)
	// 5. Cetak: fmt.Println("Backup:", backup)
	fmt.Println("Backup:", backup)

	_ = fmt.Println
	_ = original
}
