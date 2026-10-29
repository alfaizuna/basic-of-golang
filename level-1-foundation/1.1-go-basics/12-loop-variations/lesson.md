# Loop Variations

Here's your first look at advanced loop variations in Go: iterating with `for range`, nested loops, and breaking out with labels.

## Contoh
```go
// example.go
package main

import "fmt"

func main() {
	// 1. for range pada string
	// Menghasilkan indeks byte dan karakter (rune)
	for _, char := range "Go" {
		fmt.Printf("Karakter: %c\n", char)
	}

	// 2. Loop bersarang (nested loop) dengan Label
	// Label memungkinkan break keluar langsung dari loop terluar
Search:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j == 4 {
				fmt.Printf("Ditemukan pertama kali pada i=%d, j=%d\n", i, j)
				break Search // langsung keluar dari loop luar (Search)
			}
		}
	}
}
```

## Penjelasan
- `for i, v := range collection` → Mengiterasi elemen koleksi (pada string: `i` adalah indeks byte dan `v` adalah karakter / rune).
- `for _, v := range collection` → Gunakan blank identifier `_` jika tidak membutuhkan nilai indeksnya.
- Loop bersarang (*nested loop*) → Meletakkan perulangan `for` di dalam perulangan `for` lainnya.
- Label loop (misal `OuterLoop:`) → Penanda nama tepat di atas perulangan `for`.
- `break OuterLoop` → `break` biasa hanya keluar dari loop terdalam. Menambahkan nama label membuat `break` langsung menghentikan loop terluar yang ditandai.

## >_ Exercise
Buka file `main.go`:
1. Buat fungsi `countVowels(text string) int`:
   - Gunakan `for _, char := range text` untuk mengiterasi setiap karakter.
   - Hitung banyaknya huruf vokal kecil: `'a'`, `'i'`, `'u'`, `'e'`, `'o'`.
   - Kembalikan total jumlah vokal.
2. Buat fungsi `findTarget(matrixSize int, targetSum int) (int, int)`:
   - Siapkan variabel `foundR, foundC := 0, 0`.
   - Buat loop bersarang: baris `r` dari `1` s/d `matrixSize`, dan kolom `c` dari `1` s/d `matrixSize`.
   - Beri label (misalnya `SearchLoop:`) tepat di atas loop terluar `r`.
   - Jika `r + c == targetSum`:
     - Simpan `foundR = r` dan `foundC = c`.
     - Hentikan seluruh pencarian dengan `break SearchLoop`.
   - Kembalikan `foundR, foundC`.
3. Di dalam `func main()`:
   - Panggil `countVowels("golang")` dan cetak hasilnya dengan `fmt.Println`: (vokal: 'o', 'a' -> 2).
   - Panggil `findTarget(5, 7)`, simpan ke `r, c`, dan cetak: `fmt.Println(r, c)` -> (2 5).

Expected output:
```
2
2 5
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
`for range` mempermudah iterasi elemen tanpa perlu mengelola variabel counter secara manual. Pada loop bersarang, Go menyediakan *labeled break* (`break LabelName`) untuk keluar dari beberapa tingkat loop sekaligus dengan bersih tanpa perlu variabel flag tambahan.
