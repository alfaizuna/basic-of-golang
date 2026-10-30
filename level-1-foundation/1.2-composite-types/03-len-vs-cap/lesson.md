# len vs cap

Here's how to understand the difference between a slice's length and its underlying capacity.

## Contoh
```go
// example.go
//go:build ignore

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("len: %d, cap: %d\n", len(numbers), cap(numbers)) // len: 5, cap: 5

	sub := numbers[2:4]
	fmt.Printf("len: %d, cap: %d\n", len(sub), cap(sub))         // len: 2, cap: 3
}
```

## Penjelasan
- `len(s)` → mengembalikan panjang slice (jumlah elemen yang aktif terisi).
- `cap(s)` → mengembalikan kapasitas (jumlah slot dari elemen awal slice hingga ujung array di memori).
- `s[low:high]` → memotong slice dari indeks `low` sampai sebelum `high`.
- Analogi: Kapasitas (`cap`) adalah ukuran gelas, sedangkan Panjang (`len`) adalah volume air yang terisi di dalam gelas tersebut.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/03-len-vs-cap/main.go):
1. Lengkapi fungsi `getLenAndCap(s []int) (int, int)` agar mengembalikan `len(s)` dan `cap(s)`.
2. Di dalam `func main()`:
   - Buat sub-slice `part` dengan memotong `data[1:3]`.
   - Panggil `getLenAndCap(part)` dan tampung ke variabel `l, c`.
   - Cetak panjang dan kapasitasnya: `fmt.Println(l, c)`.

Expected output:
```
2 4
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
`len` menunjukkan berapa banyak elemen yang kamu miliki sekarang, sedangkan `cap` menunjukkan seberapa banyak elemen yang bisa ditampung slice sebelum Go perlu mengalokasikan memori baru di balik layar.
