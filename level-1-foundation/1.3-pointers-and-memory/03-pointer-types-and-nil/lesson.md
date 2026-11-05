# Pointer Types & nil

Here's your first look at pointer types (`*T`), their zero value (`nil`), and why Go doesn't allow pointer arithmetic.

## Contoh
example.go
```go
//go:build ignore

package main

import "fmt"

func main() {
	// 1. Deklarasi pointer tanpa nilai awal -> zero value adalah nil
	var p *int
	fmt.Println("Zero value pointer:", p) // <nil>

	// 2. Mengecek apakah pointer bernilai nil sebelum dereference
	if p == nil {
		fmt.Println("p masih nil, jangan di-dereference!")
	}

	// 3. Menghubungkan pointer ke alamat variabel yang valid
	num := 100
	p = &num
	if p != nil {
		fmt.Println("Nilai yang ditunjuk p:", *p) // 100
	}

	// Catatan: Go sengaja TIDAK mengizinkan pointer arithmetic (seperti p++ atau p + 1)
	// demi keamanan memori!
}
```

## Penjelasan
- `*int`, `*string`, `*T` → penulisan tipe data pointer di Go (pointer yang menunjuk ke data bertipe `T`).
- `nil` → *zero value* dari semua tipe pointer. Pointer yang `nil` berarti belum menunjuk ke alamat memori manapun.
- Melakukan dereference (`*p`) saat `p == nil` akan menyebabkan program berhenti mendadak (**runtime panic**).
- `if p != nil` → pola standar di Go untuk memastikan keamanan sebelum membaca atau menulis lewat pointer (*nil guard*).
- Go tidak memiliki *pointer arithmetic* (tidak bisa `ptr++` atau menggeser offset byte langsung) untuk mencegah bug memory corruption yang berbahaya seperti pada C/C++.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/03-pointer-types-and-nil/main.go):

Di dalam fungsi `main()`:
1. Deklarasikan variabel pointer bernama `ptr` dengan tipe `*string` tanpa nilai awal (biarkan zero value).
2. Periksa apakah `ptr == nil`. Jika ya, cetak:
   `fmt.Println("ptr adalah nil")`
3. Berikan alamat dari variabel `message` ke variabel `ptr` (`ptr = &message`).
4. Periksa apakah `ptr != nil`. Jika ya, cetak:
   `fmt.Println("ptr tidak nil, isinya:", *ptr)`

Expected output:
```
ptr adalah nil
ptr tidak nil, isinya: Belajar Go menyenangkan!
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Tipe pointer ditulis dengan format `*T` dan memiliki *zero value* `nil`. Selalu pastikan pointer tidak `nil` sebelum melakukan dereference untuk mencegah runtime panic, dan Go melarang *pointer arithmetic* demi keamanan memori.
