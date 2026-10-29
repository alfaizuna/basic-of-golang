# Arrays: Fixed Size

Here's your first look at Go collections: arrays with fixed sizes and value semantics.

## Contoh
```go
// example.go
//go:build ignore

package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	c := a
	c[0] = 100

	fmt.Println("a (asli)   :", a) // [1 2 3]
	fmt.Println("c (salinan):", c) // [100 2 3]
}
```

## Penjelasan
- Array di Go memiliki ukuran tetap yang ditentukan sejak awal menggunakan sintaks `[N]T` (misalnya `[3]int`).
- Panjang array adalah bagian dari tipe datanya: `[3]int` dan `[4]int` adalah dua tipe yang berbeda dan tidak kompatibel.
- Simbol `[...]` dapat digunakan pada deklarasi literal (misalnya `[...]string{"A", "B"}`) agar compiler menghitung panjang elemennya secara otomatis.
- Array di Go adalah *value type*, bukan *reference type*. Saat di-assign ke variabel lain atau dioper ke fungsi, seluruh isi array akan disalin (*copied*). Mengubah salinannya tidak akan mempengaruhi array aslinya.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/01-arrays-fixed-size/main.go):
1. Selesaikan fungsi `modifyCopy(arr [3]int) [3]int`:
   - Ubah elemen pertama: `arr[0] = 99`.
   - Kembalikan `arr`.
2. Di dalam `func main()`:
   - Kode pemanggilan dan pencetakan array sudah disiapkan.
   - Perhatikan bahwa array asli `scores` tetap bernilai `[10 20 30]` meskipun salinannya `modified` telah diubah menjadi `[99 20 30]`.

Expected output:
```
[10 20 30]
[99 20 30]
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Ukuran array di Go melekat pada tipenya dan tidak bisa bertambah atau berkurang. Karena array adalah *value type*, Go selalu menyalin seluruh isi array ketika dioper ke fungsi atau disimpan ke variabel baru.
