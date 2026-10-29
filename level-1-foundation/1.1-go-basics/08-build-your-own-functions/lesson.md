# Build Your Own Functions

Here's how to organize reusable logic using functions, parameters, and return values.

## Contoh
```go
// example.go
package main

import "fmt"

// Fungsi dengan parameter dan nilai kembalian (return value)
func add(a, b int) int {
	return a + b
}

// Fungsi tanpa return value (hanya melakukan aksi)
func sayHello(name string) {
	fmt.Println("Halo,", name)
}

func main() {
	total := add(10, 5)
	fmt.Println("Total:", total)

	sayHello("Gopher")
}
```

## Penjelasan
- `func namaFungsi(param tipe) tipeReturn { ... }` → Struktur dasar fungsi di Go.
- `(a, b int)` → Jika beberapa parameter berurutan memiliki tipe yang sama, kamu cukup menulis tipe datanya sekali di akhir.
- `return a + b` → Mengirim nilai kembali ke pemanggil. Jika fungsi mendefinisikan tipe return, wajib menyertakan kata kunci `return`.
- Posisi penulisan tipe data di Go selalu **setelah** nama parameter (contoh: `name string`, bukan `string name`).

## >_ Exercise
Buka file `main.go`:
1. Buat fungsi `greet`:
   - Menerima parameter `name string`.
   - Mengembalikan `string` dengan format `"Halo, " + name`.
2. Buat fungsi `double`:
   - Menerima parameter `n int`.
   - Mengembalikan hasil perkalian `n * 2` (bertipe `int`).
3. Di dalam `func main()`:
   - Panggil `greet("Go")` dan cetak hasilnya menggunakan `fmt.Println`.
   - Panggil `double(21)` dan cetak hasilnya menggunakan `fmt.Println`.

Expected output:
```
Halo, Go
42
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Fungsi memecah kode menjadi blok-blok kecil yang dapat digunakan kembali. Di Go, tipe data selalu ditulis setelah nama variabel/parameter, dan fungsi yang memiliki tipe return wajib mengembalikan nilai dengan kata kunci `return`.
