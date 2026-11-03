# Struct Basics

Struct adalah tipe data komposit kustom di Go yang mengelompokkan berbagai field data yang saling berhubungan menjadi satu kesatuan.

## Contoh
example.go
```go
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "Alice", Age: 28}
	u.Age = 29
	fmt.Printf("%s berumur %d tahun\n", u.Name, u.Age)
}
```

## Penjelasan
- `type NamaStruct struct { ... }` → mendefinisikan tipe data struct baru.
- `User{Name: "Alice", Age: 28}` → inisialisasi menggunakan *named field literal* (praktik terbaik di Go agar kode jelas dan tahan terhadap perubahan urutan field).
- Akses dan modifikasi field dilakukan menggunakan tanda titik `.`: `u.Name` atau `u.Age = 30`.
- Jika dibuat tanpa nilai awal (`var u User`), setiap field otomatis memiliki *zero value* dari tipe datanya masing-masing.

## >_ Exercise
Buka [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/14-struct-basics/main.go):
1. Definisikan tipe struct `Book` yang memiliki tiga field:
   - `Title` bertipe `string`
   - `Author` bertipe `string`
   - `Pages` bertipe `int`
2. Di `main()`:
   - Buat variabel `b` bertipe `Book` dengan nilai:
     - `Title: "The Go Programming Language"`
     - `Author: "Alan Donovan"`
     - `Pages: 380`
   - Cetak:
     `fmt.Printf("%s karya %s (%d halaman)\n", b.Title, b.Author, b.Pages)`

Expected output:
```
The Go Programming Language karya Alan Donovan (380 halaman)
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Struct adalah fondasi utama pembuatan model data di Go (Go tidak memiliki class). Gunakan struct untuk menggabungkan data terkait, dan biasakan memakai *named fields* saat inisialisasi.
