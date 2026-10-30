# Nil Slices & Slice Recipes

Di Go, slice kosong memiliki dua bentuk (`nil` vs `empty`), dan manipulasi elemen seperti menghapus data dilakukan menggunakan resep slicing + `append`.

## Contoh
example.go
```go
package main

import "fmt"

func main() {
	var nilSlice []int    // nilSlice == nil (true), len: 0
	emptySlice := []int{} // emptySlice == nil (false), len: 0

	letters := []string{"A", "B", "C", "D"}
	// Resep hapus indeks 1: gabungkan elemen sebelum 1 dan sesudah 1
	letters = append(letters[:1], letters[2:]...)
	fmt.Println("Setelah hapus indeks 1:", letters) // [A C D]
}
```

## Penjelasan
- `var s []T` (Nil slice) → pointer-nya `nil`, `len` 0, `cap` 0. Ini adalah cara standar & paling hemat memori untuk menginisialisasi slice kosong.
- `s := []T{}` (Empty slice) → bukan `nil` karena pointer-nya menunjuk ke memori kosong, tetapi `len` tetap 0.
- Keduanya berperilaku sama persis untuk `len(s)`, `cap(s)`, `for range`, maupun `append(s, val)`.
- **Resep Hapus Elemen pada Indeks `i`**:
  `append(s[:i], s[i+1:]...)`
  Tanda `...` (ellipsis) membongkar slice `s[i+1:]` menjadi argumen individual untuk `append`.

## >_ Exercise
Buka [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/08-nil-slices-and-slice-recipes/main.go):
1. Lengkapi fungsi `removeAt(items []string, index int) []string` untuk menghapus elemen pada `index` menggunakan pola `append(items[:index], items[index+1:]...)`.
2. Di `main()`:
   - Buat slice `languages := []string{"Go", "Java", "Python", "Ruby"}`.
   - Hapus elemen di indeks `1` (`"Java"`) menggunakan `removeAt`.
   - Cetak: `fmt.Println("Languages:", languages)`
   - Buat sebuah nil slice `var empty []string` dan cetak status nil-nya:
     `fmt.Println("Is Nil:", empty == nil)`

Expected output:
```
Languages: [Go Python Ruby]
Is Nil: true
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Nil slice (`var s []T`) adalah idiomatik di Go dan aman digunakan tanpa inisialisasi awal. Untuk memanipulasi slice (seperti menghapus elemen), manfaatkan kombinasi slicing dan `append(s[:i], s[i+1:]...)`.
