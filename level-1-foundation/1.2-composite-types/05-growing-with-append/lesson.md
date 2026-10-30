# Growing with append

Fungsi bawaan `append` memungkinkan slice bertambah panjang secara dinamis saat program berjalan.

## Contoh
example.go
```go
package main

import "fmt"

func main() {
	var numbers []int
	numbers = append(numbers, 10)
	numbers = append(numbers, 20, 30)
	fmt.Printf("len: %d, cap: %d, isi: %v\n", len(numbers), cap(numbers), numbers)
}
```

## Penjelasan
- `append(slice, elem)` → fungsi bawaan untuk menambahkan satu atau lebih elemen ke ujung slice.
- `numbers = append(numbers, 10)` → hasil `append` **harus ditampung kembali** ke variabel slice.
- Mengapa harus ditampung ulang? Jika kapasitas backing array lama habis, Go mengalokasikan backing array baru dan `append` mengembalikan slice header yang menunjuk ke lokasi baru tersebut.
- `append(numbers, 20, 30)` → kamu bisa menambahkan banyak elemen sekaligus dalam satu pemanggilan.

## >_ Exercise
Buka `main.go`. Buat fungsi helper dan lengkapi `main()`:
1. Lengkapi fungsi `addScore(scores []int, newScore int) []int` menggunakan `append` lalu kembalikan slice-nya.
2. Di `main()`:
   - Buat slice `scores` bertipe `[]int` dengan nilai awal `80` dan `90`.
   - Tambahkan `100` ke `scores` menggunakan `addScore`.
   - Tambahkan `95` ke `scores` menggunakan `addScore`.
   - Cetak: `fmt.Println("Scores:", scores)`
   - Cetak: `fmt.Println("Length:", len(scores))`

Expected output:
```
Scores: [80 90 100 95]
Length: 4
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
`append` adalah cara utama menambah elemen ke slice di Go. Karena Go mungkin memindahkan data ke backing array baru saat kapasitas penuh, selalu tampung kembali nilai kembalian dari `append`.
