# Counting & Grouping

Karakteristik zero value pada map membuat perhitungan frekuensi (counting) di Go menjadi sangat singkat dan ekspresif.

## Contoh
example.go
```go
package main

import "fmt"

func main() {
	fruits := []string{"apel", "jeruk", "apel"}
	freq := make(map[string]int)

	for _, fruit := range fruits {
		freq[fruit]++ // jika belum ada, otomatis 0 + 1
	}

	fmt.Println("Jumlah apel:", freq["apel"]) // 2
}
```

## Penjelasan
- Pola `counts[item]++` memanfaatkan zero value Go: jika `item` belum pernah ada di map, nilai default-nya adalah `0`. Menambahkan `++` langsung menjadikannya `1` tanpa perlu pengecekan `if` manual.
- Pola ini adalah salah satu idiom Go paling populer untuk menghitung frekuensi, agregasi, atau pengelompokan data.

## >_ Exercise
Buka [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/11-counting-and-grouping/main.go):
1. Lengkapi fungsi `countWords(words []string) map[string]int`:
   - Buat map hasil menggunakan `make(map[string]int)`.
   - Iterasi slice `words` dengan `for range` dan tambahkan frekuensi kemunculannya menggunakan idiom `++`.
   - Kembalikan map hasil.
2. Di `main()`:
   - Diberikan slice `words := []string{"go", "rust", "go", "python", "go", "rust"}`.
   - Panggil `counts := countWords(words)`.
   - Cetak:
     - `fmt.Println("Go:", counts["go"])`
     - `fmt.Println("Rust:", counts["rust"])`
     - `fmt.Println("Python:", counts["python"])`

Expected output:
```
Go: 3
Rust: 2
Python: 1
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Gunakan `counts[key]++` untuk menghitung frekuensi data. Karena map mengembalikan zero value untuk key baru, kamu tidak perlu menulis percabangan untuk menginisialisasi angka 0.
