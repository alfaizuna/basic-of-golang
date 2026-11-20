# Filtering & Selecting

Pipeline data sering kali hanya membutuhkan subset data tertentu (misalnya transaksi di atas nominal tertentu atau kategori tertentu saja).
Pada pelajaran keempat, kita akan membuat fungsi filter yang fleksibel menggunakan fungsi kriteria (*predicate function*), serta fungsi seleksi kategori.

## Contoh
example.go
```go
package main

type Predicate func(Row) bool

func Filter(rows []Row, keep Predicate) []Row {
	var filtered []Row
	for _, r := range rows {
		if keep(r) {
			filtered = append(filtered, r)
		}
	}
	return filtered
}
```

## Penjelasan
- `Predicate func(Row) bool`: tipe fungsi tingkat tinggi yang menerima `Row` dan mengembalikan boolean (`true` jika baris lolos seleksi).
- Pola filter berbasis predikat ini membuat kode penyaringan sangat modular dan reusable tanpa perlu membuat banyak perulangan terpisah.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.5-project-yaml-data-pipeline/04-filtering-and-selecting/main.go).

1. Definisikan tipe fungsi predikat:
   ```go
   type Predicate func(Row) bool
   ```
2. Buat fungsi umum `Filter(rows []Row, predicate Predicate) []Row`:
   - Iterasi slice `rows`.
   - Jika `predicate(r)` menghasilkan `true`, tambahkan baris tersebut ke slice hasil filter.
   - Kembalikan slice hasil filter.
3. Buat dua fungsi filter spesifik:
   - `FilterByCategory(rows []Row, category string) []Row`: menyaring baris yang memiliki `r.Category == category`.
   - `FilterMinTotal(rows []Row, minTotal float64) []Row`: menyaring baris dengan `r.Total >= minTotal`.

Expected output ketika `go run main.go`:
```
Filtered category 'electronics': 2 rows
Filtered min total $500: 2 rows
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Menggunakan fungsi sebagai argumen (`Predicate func(Row) bool`) adalah pola standar pemrograman fungsional di Go untuk membangun filter data yang bersih dan modular.
