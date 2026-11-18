# Guarding with nil Checks

Mari pelajari pola idiomatik Go untuk mencegah panic menggunakan nil guard dan return `(*T, error)`.

## Contoh
example.go
```go
package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID    int
	Title string
}

func FindProduct(id int) (*Product, error) {
	if id != 1 {
		return nil, errors.New("product not found")
	}
	return &Product{ID: 1, Title: "Keyboard"}, nil
}

func GetTitle(p *Product) string {
	if p == nil {
		return "Unknown"
	}
	return p.Title
}
```

## Penjelasan
- `if p == nil { return "Unknown" }` → guard clause untuk mengamankan fungsi sebelum dereferencing pointer.
- `return nil, errors.New(...)` → pola idiomatik fungsi yang gagal: kembalikan `nil` untuk pointer dan objek error.
- `item, err := FindProduct(...)` lalu `if err != nil` → selalu periksa error terlebih dahulu sebelum menyentuh nilai pointer kembalian.
- Hindari mengabaikan error (`item, _ := FindProduct(...)`) karena jika terjadi kegagalan, `item` bernilai `nil` dan mengaksesnya akan memicu panic.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/11-guarding-with-nil-checks/main.go).

Lengkapi dua fungsi berikut:
1. Di `FindUser`: jika ID tidak ditemukan setelah loop selesai, kembalikan `nil` dan error dengan pesan `"user not found"` menggunakan `errors.New("user not found")`.
2. Di `FormatUserName`: tambahkan *guard check* di awal fungsi. Jika `u == nil`, kembalikan string `"Guest"`. Jika tidak, kembalikan `u.Name`.

Expected output:
```
User: Alice
Error: user not found
Fallback: Guest
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Selalu periksa `err != nil` sebelum menggunakan pointer hasil kembalian sebuah fungsi, dan gunakan *nil guard clause* (`if ptr == nil`) di dalam fungsi yang menerima pointer opsional agar terhindar dari runtime panic.
