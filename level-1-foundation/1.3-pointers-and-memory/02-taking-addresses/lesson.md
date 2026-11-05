# Taking Addresses: & and *

Here's your first look at how to get a variable's memory address using `&` and access or modify its value using `*`.

## Contoh
example.go
```go
//go:build ignore

package main

import "fmt"

func main() {
	age := 25

	// 1. Mengambil alamat memori dengan operator & (address-of)
	ptr := &age
	fmt.Println("Alamat memori age:", ptr) // Contoh output: 0x14000122010

	// 2. Membaca nilai dari alamat tersebut dengan operator * (dereference)
	fmt.Println("Nilai age lewat pointer:", *ptr) // 25

	// 3. Mengubah nilai langsung di alamat memori tersebut
	*ptr = 30
	fmt.Println("Nilai age setelah diubah:", age) // 30
}
```

## Penjelasan
- `&age` → operator *address-of* (`&`) menghasilkan alamat memori tempat variabel `age` disimpan.
- `ptr := &age` → variabel `ptr` bertipe pointer ke integer (`*int`) yang menyimpan alamat memori `age`.
- `*ptr` → operator *dereference* (`*`) mengakses nilai asli yang berada di alamat yang ditunjuk oleh `ptr`.
- `*ptr = 30` → mengubah nilai di alamat memori tersebut secara langsung; variabel `age` ikut terupdate tanpa perlu di-assign ulang lewat `age = 30`.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/02-taking-addresses/main.go):
Di dalam fungsi `main()`:
1. Ambil alamat memori dari `score` menggunakan operator `&` dan simpan ke variabel baru bernama `ptr`.
2. Cetak nilai `score` melalui pointer `ptr` (dereference `*ptr`) dengan tulisan:
   `fmt.Println("Nilai awal lewat pointer:", *ptr)`
3. Ubah nilai `score` menjadi `99` langsung melalui pointer `ptr` menggunakan `*ptr = 99`.
4. Cetak nilai `score` asli:
   `fmt.Println("Nilai score setelah diubah:", score)`

Expected output:
```
Nilai awal lewat pointer: 42
Nilai score setelah diubah: 99
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Operator `&` digunakan untuk mengambil alamat memori suatu variabel (*address-of*), sedangkan operator `*` digunakan untuk membaca atau mengubah nilai di alamat tersebut (*dereference*).
