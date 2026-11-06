# new and &T{}

Here's your first look at creating pointers dynamically using the `new()` function and struct literal addresses (`&T{}`).

## Contoh
example.go
```go
//go:build ignore

package main

import "fmt"

type Database struct {
	Name    string
	Timeout int
}

func main() {
	// 1. Menggunakan new() untuk tipe dasar
	// new(T) mengalokasikan memori untuk T, mengisinya dengan zero value, dan mengembalikan *T
	pInt := new(int)
	fmt.Println("Zero value via new(int):", *pInt) // 0
	*pInt = 42
	fmt.Println("Setelah diubah:", *pInt)          // 42

	// 2. Menggunakan new() untuk struct
	// Semua field otomatis bernilai zero value ("" dan 0)
	db1 := new(Database)
	fmt.Println("db1 via new():", db1) // &{ 0}

	// 3. Menggunakan &T{} (struct literal) — cara paling idiomatik di Go!
	// Mengalokasikan struct sekaligus menginisialisasi field yang diinginkan
	db2 := &Database{
		Name:    "postgres",
		Timeout: 30,
	}
	fmt.Println("db2 via &T{}:", db2) // &{postgres 30}
}
```

## Penjelasan
- `new(T)` → fungsi bawaan Go yang mengalokasikan memori seukuran tipe `T`, mengisinya dengan *zero value* dari `T`, lalu mengembalikan pointer bertipe `*T`.
- `new(int)` menghasilkan `*int` bernilai `0`. `new(Database)` menghasilkan `*Database` dengan semua field berupa *zero value*.
- `&T{...}` → mengalokasikan struct literal dan langsung mengambil alamat memorinya.
- Kapan memakai yang mana?
  - `new(T)` biasanya digunakan untuk tipe primitif bila membutuhkan pointer ke nilai nol (meski jarang).
  - `&T{...}` adalah cara standar dan paling idiomatik di Go untuk struct karena memungkinkan kita menginisialisasi nilai awal field dalam satu langkah ekspresif.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/04-new-and-struct-literals/main.go):

1. Di dalam `main()`, alokasikan pointer integer menggunakan `new(int)` dan simpan ke variabel `version`.
2. Ubah nilai yang ditunjuk oleh `version` menjadi `1` menggunakan operator dereference (`*version = 1`).
3. Cetak nilai `version`:
   ```go
   fmt.Println("Version:", *version)
   ```
4. Buat pointer ke struct `ServerConfig` menggunakan sintaks `&ServerConfig{...}` dengan field:
   - `Host: "localhost"`
   - `Port: 8080`
   Simpan ke variabel `cfg`.
5. Cetak konfigurasi server menggunakan `fmt.Printf`:
   ```go
   fmt.Printf("Server running at %s:%d\n", cfg.Host, cfg.Port)
   ```

Expected output:
```
Version: 1
Server running at localhost:8080
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Gunakan `new(T)` untuk mengalokasikan memori berukuran tipe `T` yang berisi *zero value* dan mengembalikan pointernya. Untuk struct, idiom Go lebih menyukai `&T{...}` karena memungkinkan inisialisasi field secara langsung dan bersih.
