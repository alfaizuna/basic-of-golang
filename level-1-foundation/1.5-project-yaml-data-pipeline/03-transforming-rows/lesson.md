# Transforming Rows

Setelah data mentah berhasil dimuat ke dalam slice `[]Row`, langkah penting berikutnya dalam data pipeline adalah **transformasi data**:
1. Menghitung field turunan yang belum terisi (misalnya `Total = Price * float64(Quantity)`).
2. Menormalisasi format teks (misalnya mengubah `Category` menjadi huruf kecil dan membuang spasi liar di awal/akhir).

## Contoh
example.go
```go
package main

import "strings"

func CalculateTotals(rows []Row) []Row {
	result := make([]Row, len(rows))
	copy(result, rows)

	for i := range result {
		result[i].Total = result[i].Price * float64(result[i].Quantity)
	}
	return result
}
```

## Penjelasan
- Membuat slice baru dengan `copy(result, rows)` mencegah efek samping (*side effects*) yang merusak data input asli.
- Iterasi menggunakan indeks `result[i]` wajib digunakan agar perubahan field struct benar-benar tersimpan di backing array, bukan sekadar memodifikasi variabel salinan perulangan.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.5-project-yaml-data-pipeline/03-transforming-rows/main.go).

Buat dua fungsi transformasi:
1. `CalculateTotals(rows []Row) []Row`:
   - Buat salinan slice `rows`.
   - Untuk setiap baris, hitung `Total = Price * float64(Quantity)`.
   - Kembalikan slice hasil kalkulasi.
2. `NormalizeCategories(rows []Row) []Row`:
   - Buat salinan slice `rows`.
   - Untuk setiap baris, ubah field `Category` menjadi huruf kecil dan pangkas spasi menggunakan `strings.ToLower(strings.TrimSpace(r.Category))`.
   - Kembalikan slice hasil normalisasi.

Expected output ketika `go run main.go`:
```
Transformed: TX-101 | Laptop | electronics | Total: $2400.00
Transformed: TX-102 | Desk | furniture | Total: $450.00
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Transformasi baris dilakukan dengan mengiterasi slice lewat indeks (`rows[i]`) agar mutasi field tercatat langsung pada elemen slice.
