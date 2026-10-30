# Copy: Independent Slice Copies

Fungsi bawaan `copy` menduplikasi elemen antar slice sehingga masing-masing memiliki backing array independen.

## Contoh
example.go
```go
package main

import "fmt"

func main() {
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)

	dst[0] = 999
	fmt.Println("src:", src) // [1 2 3] (tidak berubah!)
	fmt.Println("dst:", dst) // [999 2 3]
}
```

## Penjelasan
- `copy(dst, src)` → menyalin elemen dari slice `src` ke slice `dst`. Mengembalikan jumlah elemen yang berhasil disalin.
- **Penting:** `dst` **harus sudah memiliki alokasi `len`**. Jumlah elemen yang disalin adalah `min(len(dst), len(src))`. Jika `dst` masih kosong (`len: 0`), tidak ada data yang tersalin!
- Karena `dst` dialokasikan secara terpisah (misalnya lewat `make`), modifikasi pada `dst` **tidak mempengaruhi** `src`.

## >_ Exercise
Buka [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/07-copy/main.go):
1. Diberikan slice awal `original := []string{"Go", "Rust", "Python"}`.
2. Buat slice `backup` menggunakan `make` dengan panjang (`len`) yang sama dengan `original`.
3. Salin seluruh isi `original` ke `backup` menggunakan `copy`.
4. Ubah elemen terakhir pada `backup` menjadi `"TypeScript"`.
5. Cetak:
   - `fmt.Println("Original:", original)`
   - `fmt.Println("Backup:", backup)`

Expected output:
```
Original: [Go Rust Python]
Backup: [Go Rust TypeScript]
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
`copy(dst, src)` menyalin data sebanyak `min(len(dst), len(src))`. Pastikan slice tujuan (`dst`) sudah dialokasikan dengan panjang yang cukup agar data benar-benar tersalin secara independen.
