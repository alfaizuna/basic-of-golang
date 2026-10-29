# Making Decisions

Here's how to control program flow using `if`, `else`, and Go's clean initializer syntax.

## Contoh
```go
// example.go
package main

import "fmt"

func main() {
	// 1. If-else biasa (tanpa tanda kurung pada kondisi)
	temperature := 30
	if temperature > 35 {
		fmt.Println("Panas sekali")
	} else if temperature > 25 {
		fmt.Println("Hangat nyaman")
	} else {
		fmt.Println("Sejuk")
	}

	// 2. If dengan initializer: deklarasi variabel langsung sebelum kondisi
	// Variabel 'status' hanya hidup di dalam blok if-else ini
	if status := 200; status == 200 {
		fmt.Println("OK")
	}
}
```

## Penjelasan
- `if temperature > 35` → Di Go, kondisi `if` **tidak perlu** diapit tanda kurung `()`.
- `} else {` → Kurung kurawal pembuka `{` dan penutup `}` wajib ada. Kata kunci `else` wajib berada di baris yang sama dengan `}` penutup sebelumnya (`} else {`), jika tidak Go akan memicu syntax error.
- `if status := 200; status == 200` → *If with initializer*. Kamu bisa membuat variabel singkat sebelum tanda titik koma `;`. Variabel tersebut hanya aktif di dalam blok keputusan itu saja (*scoped*).

## >_ Exercise
Buka file `main.go`. Buatlah penentu nilai ujian (*grade calculator*) menggunakan `if` dengan initializer untuk mengevaluasi skor `score := 85`:
- Jika `score >= 90` → cetak `"Grade: A"`
- Jika `score >= 80` → cetak `"Grade: B"`
- Jika `score >= 70` → cetak `"Grade: C"`
- Selain itu → cetak `"Grade: F"`

Expected output:
```
Grade: B
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Kondisi `if` di Go bersih tanpa kurung `()`, dan format `} else {` wajib dalam satu baris. Gunakan *if dengan initializer* untuk menjaga kode tetap rapi dan membatasi variabel hanya di tempat ia dibutuhkan.
