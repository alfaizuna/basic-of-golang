# Slice Basics

Here's your first look at Go slices: dynamic collections built on top of arrays.

## Contoh
```go
// example.go
//go:build ignore

package main

import "fmt"

func main() {
	fruits := []string{"Apel", "Jeruk", "Mangga"}

	fmt.Printf("Panjang: %d, Elemen pertama: %s\n", len(fruits), fruits[0])

	for i, fruit := range fruits {
		fmt.Printf("[%d] %s\n", i, fruit)
	}
}
```

## Penjelasan
- `[]T{...}` → slice literal; dideklarasikan tanpa angka ukuran di dalam kurung siku `[]`, berbeda dari array `[N]T`.
- `len(s)` → mengembalikan jumlah elemen yang ada di dalam slice saat ini.
- `s[i]` → mengakses elemen pada indeks ke-`i` (dimulai dari indeks `0`).
- `for i, v := range s` → mengiterasi slice untuk membaca indeks `i` dan nilai elemen `v`.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/02-slice-basics/main.go):
1. Selesaikan fungsi `sumSlice(numbers []int) int`:
   - Gunakan perulangan `for _, num := range numbers` untuk menjumlahkan semua angka di dalam slice.
   - Kembalikan total penjumlahannya.
2. Di dalam `func main()`:
   - Slice `prices := []int{15, 25, 35, 45}` dan pencetakan jumlah elemen serta hasil penjumlahan sudah disiapkan.
   - Lengkapi fungsi `sumSlice` agar program menghasilkan output yang diharapkan.

Expected output:
```
Count: 4
Total: 120
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Slice (`[]T`) adalah tipe data koleksi paling umum di Go. Tidak seperti array yang berukuran kaku, slice bersifat fleksibel dan menyediakan akses elemen serta iterasi yang sangat praktis melalui `len()` dan `for range`.
