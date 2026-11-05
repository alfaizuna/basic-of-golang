# Values Are Copies

Here's your first look at how Go handles data when passing it to functions: by default, everything in Go is passed by value (copied).

## Contoh
example.go
```go
//go:build ignore

package main

import "fmt"

func double(n int) {
	n = n * 2
	fmt.Println("Di dalam fungsi double:", n) // 20
}

func main() {
	x := 10
	double(x)
	fmt.Println("Di dalam main:", x) // tetap 10!
}
```

## Penjelasan
- `double(x)` → Go menyalin nilai `10` dari variabel `x` ke variabel parameter lokal baru bernama `n`.
- `n = n * 2` → perubahan hanya terjadi pada salinan lokal di dalam fungsi `double`.
- `fmt.Println("Di dalam main:", x)` → nilai variabel `x` pada fungsi `main` tidak berubah sama sekali karena berada di lokasi memori yang berbeda.
- Sifat ini disebut **pass by value**. Jika kita ingin fungsi lain dapat mengubah data asli pemanggilnya, kita membutuhkan **pointer** (yang akan dipelajari di lesson berikutnya).

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/01-values-are-copies/main.go):
1. Lengkapi fungsi `resetCount(count int)`:
   - Ubah nilai parameter: `count = 0`
   - Cetak pesan: `fmt.Println("Di dalam resetCount:", count)`
2. Di dalam `func main()`:
   - Panggil `resetCount(total)`
   - Cetak nilai variabel asli: `fmt.Println("Di dalam main:", total)`

Expected output:
```
Di dalam resetCount: 0
Di dalam main: 100
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Di Go, semua argumen yang dikirim ke fungsi selalu disalin (*pass by value*). Mengubah nilai parameter di dalam fungsi tidak akan pernah mempengaruhi variabel asli milik si pemanggil.
