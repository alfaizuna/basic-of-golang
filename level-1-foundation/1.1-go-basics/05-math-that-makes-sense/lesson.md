# Math That Makes Sense

Here's how arithmetic and type conversions work in Go.

## Contoh
```go
// example.go
package main

import "fmt"

func main() {
	// Pembagian integer vs pembagian float
	intDiv := 7 / 2       // Hasilnya 3 (bagian desimal dibuang)
	mod := 7 % 2          // Hasilnya 1 (sisa bagi)
	floatDiv := 7.0 / 2.0 // Hasilnya 3.5

	// Konversi tipe eksplisit
	var a int = 10
	var b float64 = 2.5
	// total := a + b // ERROR! mismatched types int and float64
	total := float64(a) + b // Valid: 12.5

	fmt.Println(intDiv)
	fmt.Println(mod)
	fmt.Println(floatDiv)
	fmt.Println(total)
}
```

## Penjelasan
- `7 / 2` → Pembagian antar integer membuang pecahan (desimal), sehingga `7 / 2` menghasilkan `3`.
- `7 % 2` → Operator modulus (`%`) mengembalikan sisa pembagian (7 dibagi 2 bersisa `1`).
- `float64(a)` → Konversi tipe eksplisit. Go **tidak pernah** mengubah tipe data secara otomatis (*no implicit conversion*). Operasi antara `int` dan `float64` akan memicu compile error jika tidak dikonversi secara eksplisit.

## >_ Exercise
Buka file `main.go`. Diberikan data awal:
- `totalItems := 7`
- `bags := 2`
- `price := 12.5`
- `qty := 4`

Lengkapi perhitungan berikut:
1. `itemsPerBag`: hitung pembagian integer antara `totalItems` dan `bags`.
2. `remainder`: hitung sisa bagi `totalItems` dibagi `bags` menggunakan operator modulus `%`.
3. `avgItems`: konversikan `totalItems` dan `bags` ke `float64`, lalu hitung hasil pembagian desimalnya.
4. `totalPrice`: kalikan `price` dengan `qty` (konversikan `qty` ke `float64` terlebih dahulu).
5. Cetak berturut-turut di baris terpisah: `itemsPerBag`, `remainder`, `avgItems`, dan `totalPrice`.

Expected output:
```
3
1
3.5
50
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Di Go, pembagian antar integer selalu menghasilkan integer (pecahan dibuang). Go menolak konversi tipe otomatis, jadi selalu ubah tipe data secara eksplisit menggunakan sintaks `T(nilai)`.
