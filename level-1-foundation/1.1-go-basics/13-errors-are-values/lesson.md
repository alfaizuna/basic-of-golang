# Errors Are Values

Here's how Go handles errors as regular values rather than exceptions.

## Contoh
```go
// example.go
//go:build ignore

package main

import (
	"errors"
	"fmt"
)

func checkAge(age int) (string, error) {
	if age < 0 {
		return "", errors.New("age cannot be negative")
	}
	return "Valid", nil
}

func main() {
	status, err := checkAge(-5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Status:", status)
}
```

## Penjelasan
- Di Go, error bukanlah *exception* (tidak ada konsep `try-catch`), melainkan nilai biasa bertipe `error`.
- Fungsi yang berpotensi gagal mengembalikan `error` di posisi nilai balik (*return*) terakhir: `(T, error)`.
- `errors.New("pesan")` dari paket standard `"errors"` digunakan untuk membuat nilai error baru.
- Nilai `nil` pada error menandakan bahwa operasi berhasil tanpa error.
- Pola idiomatik pengecekan error di Go adalah `if err != nil` tepat setelah fungsi dipanggil.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.1-go-basics/13-errors-are-values/main.go):
1. Selesaikan fungsi `divide(a, b float64) (float64, error)`:
   - Jika `b == 0`, kembalikan `0` dan `errors.New("cannot divide by zero")`.
   - Jika tidak, kembalikan hasil `a / b` dan `nil`.
2. Di dalam `func main()`:
   - Pemanggilan pertama `divide(10, 2)` sudah disiapkan. Setelah kamu memperbaiki fungsi `divide`, ia akan mencetak `5`.
   - Buat pemanggilan kedua `divide(10, 0)`. Tangani error menggunakan pola `if err2 != nil`: jika terjadi error cetak `err2`, jika sukses cetak `res2`.

Expected output:
```
5
cannot divide by zero
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Di Go, error diperlakukan sebagai nilai biasa (*errors are values*). Tidak ada `try-catch`; setiap potensi error ditangani secara eksplisit menggunakan pola `if err != nil` di tempat terjadinya.
