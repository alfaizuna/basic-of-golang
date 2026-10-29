# Bytes, Runes & Sized Types

Here's how Go represents raw bytes, Unicode characters, and fixed-size integers.

## Contoh
```go
// example.go
//go:build ignore

package main

import "fmt"

func main() {
	var b byte = 'A' // alias uint8 (0 s/d 255)
	fmt.Printf("byte: %d (%c)\n", b, b)

	var r rune = '🚀' // alias int32 (Unicode Code Point)
	fmt.Printf("rune: %d (%c)\n", r, r)

	var maxInt8 int8 = 127
	var nextInt8 int8 = maxInt8 + 1 // overflow ke -128
	fmt.Printf("int8: %d + 1 = %d\n", maxInt8, nextInt8)
}
```

## Penjelasan
- `byte` adalah alias untuk `uint8` (1 byte / 8 bit, nilai 0–255), cocok untuk data biner mentah atau karakter ASCII.
- `rune` adalah alias untuk `int32` (4 byte / 32 bit), digunakan untuk merepresentasikan satu Unicode Code Point (mendukung karakter semua bahasa & emoji).
- Tipe angka berukuran tetap (`int8`, `int16`, `int32`, `int64`) memiliki batas memori yang pasti (contoh `int8` bernilai -128 hingga 127).
- *Integer Overflow*: Saat nilai tipe integer berukuran tetap melewati batas maksimumnya saat runtime, nilainya akan berputar (*wrap around*) ke nilai minimumnya tanpa memicu error atau *panic*.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.1-go-basics/14-bytes-runes-and-sized-types/main.go):
1. Selesaikan fungsi `toByte(r rune) (byte, bool)`:
   - Jika `r <= 255`, kembalikan `byte(r), true`.
   - Jika `r > 255`, kembalikan `0, false`.
2. Selesaikan fungsi `addInt8(a int8, b int8) int8`:
   - Kembalikan hasil `a + b` (tipe `int8`).
3. Di dalam `func main()`:
   - Panggilan pertama `toByte('A')` sudah disediakan.
   - Panggil `toByte('世')`, simpan ke `b2, ok2`, dan cetak: `fmt.Println(b2, ok2)`.
   - Panggil `addInt8(127, 1)` dan cetak hasilnya dengan `fmt.Println`.

Expected output:
```
65 true
0 false
-128
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
`byte` (uint8) digunakan untuk data mentah 1-byte, sedangkan `rune` (int32) mewakili karakter Unicode universal. Tipe bilangan berukuran tetap seperti `int8` akan mengalami *wrap-around overflow* saat perhitungannya melebihi batas nilai maksimum tipe tersebut.
