# The Building Blocks

Here's your first look at Go's fundamental data types.

## Contoh
```go
// example.go
package main

import "fmt"

func main() {
	fmt.Println(42)           // int
	fmt.Println(3.14)         // float64
	fmt.Println("Belajar Go") // string
	fmt.Println(true)         // bool
}
```

## Penjelasan
- `int` → Bilangan bulat tanpa koma (misal: `42`, `-7`).
- `float64` → Bilangan desimal/pecahan menggunakan tanda titik `.` (misal: `3.14`, `0.5`).
- `string` → Rangkaian teks yang selalu diapit tanda kutip ganda `""` (misal: `"Belajar Go"`).
- `bool` → Tipe boolean yang hanya bernilai `true` (benar) atau `false` (salah).

## >_ Exercise
Buka file `main.go`, lalu cetak data profil berikut baris per baris di dalam fungsi `main()`:
1. Angka integer: `25`
2. Angka float: `175.5`
3. Teks string: `Go`
4. Nilai boolean: `true`

Expected output:
```
25
175.5
Go
true
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Empat tipe data pondasi Go adalah `int` (bilangan bulat), `float64` (desimal), `string` (teks), dan `bool` (benar/salah). Tipe data menentukan jenis informasi yang disimpan dan operasi apa yang boleh dilakukan.
