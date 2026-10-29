# Slug Generator 4: Ship It

Here's the final step: packaging all stages into a robust `GenerateSlug` function with proper error handling and final end-to-end tests.

## Contoh
```go
// example.go
//go:build ignore

package main

import (
	"errors"
	"fmt"
	"strings"
)

func validateUsername(name string) (string, error) {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return "", errors.New("username cannot be empty")
	}
	return strings.ToLower(trimmed), nil
}

func main() {
	user, err := validateUsername("  Alfaizuna  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Username valid: %q\n", user)

	_, err2 := validateUsername("   ")
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
}
```

## Penjelasan
- Di tahap akhir ini, kita merangkum seluruh logika ke dalam fungsi publik `GenerateSlug(title string) (string, error)`.
- Menggunakan pola *guard clause* (pengecekan awal):
  - Jika input kosong setelah di-`TrimSpace`, kembalikan error `"input text cannot be empty"`.
  - Jika setelah ditransformasi menghasilkan slug kosong (misal input hanya berisi simbol tak valid), kembalikan error `"no valid slug characters found"`.
  - Jika berhasil, kembalikan slug dan `nil`.
- Pola idiomatik Go: selalu periksa `if err != nil` sebelum menggunakan hasil fungsi.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.1-go-basics/slug-generator/main.go):
1. Tambahkan paket `"errors"` pada deklarasi `import` di bagian atas.
2. Buat fungsi baru `GenerateSlug(title string) (string, error)`:
   - Jika `strings.TrimSpace(title) == ""`:
     - Kembalikan `"", errors.New("input text cannot be empty")`.
   - Jalankan `slug := transformToSlug(title)`.
   - Jika `slug == ""`:
     - Kembalikan `"", errors.New("no valid slug characters found")`.
   - Jika valid, kembalikan `slug, nil`.
3. Di dalam `func main()`:
   - Ganti isi fungsi `main()` dengan:
     ```go
     title1 := "  Belajar Go: Dasar #1 (Pemula)!  "
     slug1, err1 := GenerateSlug(title1)
     if err1 != nil {
         fmt.Println("Error:", err1)
     } else {
         fmt.Printf("Input: %q\nSlug:  %q\n", title1, slug1)
     }

     title2 := "### ??? !!!"
     _, err2 := GenerateSlug(title2)
     if err2 != nil {
         fmt.Println("Error:", err2)
     }
     ```

Expected output:
```
Input: "  Belajar Go: Dasar #1 (Pemula)!  "
Slug:  "belajar-go-dasar-1-pemula"
Error: no valid slug characters found
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Selamat! Kamu telah membangun aplikasi pembuat slug lengkap dari awal. Menggabungkan fungsi modular dengan penanganan error eksplisit (`error as value`) adalah ciri khas kode Go yang tangguh dan siap produksi (*idiomatic & production-ready*).
