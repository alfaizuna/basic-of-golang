# Constructor Functions

Karena Go tidak memiliki kata kunci khusus seperti `constructor` di bahasa lain, konvensi standar Go adalah membuat fungsi helper bernama `New<Tipe>` atau `New` untuk instansiasi dan validasi.

## Contoh
example.go
```go
package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
}

func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, errors.New("nama tidak boleh kosong")
	}
	return &User{Name: name}, nil
}
```

## Penjelasan
- Konvensi penamaan: `New<Tipe>` (misal `NewUser`, `NewProduct`) atau cukup `New` jika berada di dalam package tersendiri (misal `user.New()`).
- Fungsi konstruktor biasanya mengembalikan pointer ke struct (`*T`) dan `error`.
- Keuntungan fungsi konstruktor:
  1. Menjamin validasi data (mencegah struct dibuat dalam state yang rusak/invalid).
  2. Memberikan nilai default untuk field tertentu.

## >_ Exercise
Buka [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/17-constructor-functions/main.go):
1. Diberikan struct:
   `type Product struct { Name string; Price float64; Stock int }`
2. Lengkapi fungsi konstruktor `NewProduct(name string, price float64) (*Product, error)`:
   - Jika `name == ""`, return `nil, errors.New("nama produk wajib diisi")`
   - Jika `price <= 0`, return `nil, errors.New("harga harus lebih besar dari 0")`
   - Jika valid, kembalikan pointer `&Product{Name: name, Price: price, Stock: 0}, nil`
3. Di `main()`:
   - Buat produk valid: `p, err := NewProduct("Laptop", 12000000)`
   - Jika `err == nil`, cetak: `fmt.Printf("Produk: %s, Harga: Rp%.0f, Stok: %d\n", p.Name, p.Price, p.Stock)`
   - Coba buat produk invalid: `_, errInvalid := NewProduct("", 5000)`
   - Cetak: `fmt.Println("Error:", errInvalid)`

Expected output:
```
Produk: Laptop, Harga: Rp12000000, Stok: 0
Error: nama produk wajib diisi
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Gunakan pola `New<Tipe>` untuk menginisialisasi struct yang memerlukan validasi atau nilai default. Mengembalikan `(*T, error)` adalah idiomatik di Go.
