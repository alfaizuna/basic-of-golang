# Creating Slices with make

Here's how to allocate slices with specific lengths and capacities using the built-in `make` function.

## Contoh
```go
// example.go
//go:build ignore

package main

import "fmt"

func main() {
	zeros := make([]int, 3)        // len: 3, cap: 3, isi: [0 0 0]
	buffer := make([]string, 1, 4) // len: 1, cap: 4, isi: [""]

	buffer[0] = "Ready"
	fmt.Printf("len: %d, cap: %d, isi: %v\n", len(buffer), cap(buffer), buffer)
}
```

## Penjelasan
- `make([]T, len)` → membuat slice dengan panjang `len` (elemen terisi *zero value*) dan kapasitas awal otomatis sama dengan `len`.
- `make([]T, len, cap)` → membuat slice dengan panjang `len` dan cadangan kapasitas memori sebesar `cap`.
- `s[i] = val` → elemen hanya bisa diakses via indeks jika indeks tersebut berada dalam rentang `0` sampai `len-1` (bukan `cap-1`).
- Kapan pakai `make` vs `var`:
  - `var s []int` → *nil slice* (`len = 0, cap = 0`), ideal jika slice belum tentu terpakai.
  - `make([]T, 0, cap)` → *pre-allocated slice*, ideal jika kamu sudah tahu estimasi jumlah data yang akan ditambahkan.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/04-creating-slices-with-make/main.go):
1. Selesaikan fungsi `createBufferedSlice(length int, capacity int) []int` agar membuat slice menggunakan `make([]int, length, capacity)`.
2. Di dalam `func main()`:
   - Panggil `s := createBufferedSlice(2, 5)`.
   - Isi elemen ke-0 dengan `100` dan elemen ke-1 dengan `200`.
   - Cetak isi slice: `fmt.Println(s)`.
   - Cetak panjang dan kapasitasnya: `fmt.Println(len(s), cap(s))`.

Expected output:
```
[100 200]
2 5
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Gunakan `make([]T, len, cap)` untuk membuat slice dengan kontrol penuh atas ukuran awal dan kapasitas cadangannya. Mengalokasikan kapasitas di awal mencegah Go melakukan alokasi ulang memori yang boros di kemudian hari.
