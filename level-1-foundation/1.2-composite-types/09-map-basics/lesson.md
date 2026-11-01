# Map Basics

Map adalah struktur data bawaan Go untuk menyimpan pasangan kunci-nilai (key-value) dengan akses yang sangat cepat.

## Contoh
example.go
```go
package main

import "fmt"

func main() {
	prices := map[string]int{
		"Kopi": 15000,
		"Teh":  10000,
	}
	prices["Roti"] = 12000
	delete(prices, "Teh")

	fmt.Println("Kopi:", prices["Kopi"])
	fmt.Println("Total item:", len(prices))
}
```

## Penjelasan
- `map[K]V` → tipe data map dengan tipe kunci `K` dan tipe nilai `V`.
- `make(map[K]V)` atau `map[K]V{...}` → cara membuat dan menginisialisasi map.
- `m[key] = value` → menyimpan atau memperbarui nilai.
- `delete(m, key)` → fungsi bawaan untuk menghapus key dari map.
- `len(m)` → mengembalikan jumlah pasangan key-value yang tersimpan.

## >_ Exercise
Buka [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/09-map-basics/main.go):
1. Buat map `scores` bertipe `map[string]int` dengan nilai awal: `"Alice": 90` dan `"Bob": 85`.
2. Tambahkan entri baru: `"Charlie": 95`.
3. Hapus entri `"Bob"` menggunakan fungsi `delete`.
4. Cetak:
   - `fmt.Println("Alice:", scores["Alice"])`
   - `fmt.Println("Total:", len(scores))`

Expected output:
```
Alice: 90
Total: 2
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Map adalah tipe referensi key-value di Go. Selalu inisialisasi map sebelum menulis (menggunakan `make` atau literal `{}`), dan gunakan fungsi bawaan `delete(m, key)` untuk menghapus elemen.
