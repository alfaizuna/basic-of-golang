# Repeating Things

Here's your first look at Go's only looping keyword, `for`, and its three powerful forms.

## Contoh
```go
// example.go
package main

import "fmt"

func main() {
	// Bentuk 1: Klasik (init; condition; post)
	for i := 1; i <= 3; i++ {
		fmt.Println("Klasik:", i)
	}

	// Bentuk 2: Mirip "while" (hanya kondisi)
	count := 1
	for count <= 3 {
		fmt.Println("While-style:", count)
		count++
	}

	// Bentuk 3: Infinite loop dengan break dan continue
	num := 0
	for {
		num++
		if num == 2 {
			continue // lewati angka 2
		}
		if num > 4 {
			break // berhenti saat num > 4
		}
		fmt.Println("Infinite loop:", num)
	}
}
```

## Penjelasan
- Di Go **hanya ada satu kata kunci perulangan**: `for` (tidak ada `while` atau `do-while`).
- `for i := 1; i <= n; i++` → Bentuk standar dengan inisialisasi counter, kondisi berhenti, dan post-statement.
- `for count <= n` → Bentuk kondisi tunggal, berfungsi persis seperti perulangan `while` di bahasa lain.
- `for { ... }` → Loop tak hingga (*infinite loop*), biasanya dihentikan dari dalam menggunakan kondisi dan `break`.
- `continue` → Melompati sisa baris kode pada iterasi saat ini dan langsung maju ke iterasi berikutnya.
- `break` → Menghentikan perulangan seketika dan langsung keluar dari blok `for`.

## >_ Exercise
Buka file `main.go`:
1. Buat fungsi `sumEven(limit int) int`:
   - Gunakan perulangan `for` klasik dari `1` sampai `limit` (inklusif: `i <= limit`).
   - Jika angka ganjil (`i % 2 != 0`), gunakan `continue` untuk melompatinya.
   - Tambahkan hanya angka genap ke variabel `total`, lalu kembalikan `total`.
2. Buat fungsi `countUntil(target int) int`:
   - Mulai dengan `count := 0` dan `steps := 0`.
   - Gunakan `for { ... }` (*infinite loop*).
   - Di setiap iterasi:
     - Tambah `count` sebanyak 3 (`count += 3`).
     - Tambah `steps` sebanyak 1 (`steps++`).
     - Jika `count >= target`, hentikan perulangan dengan `break`.
   - Kembalikan `steps`.
3. Di dalam `func main()`:
   - Panggil `sumEven(10)` dan cetak hasilnya dengan `fmt.Println`.
   - Panggil `countUntil(10)` dan cetak hasilnya dengan `fmt.Println`.

Expected output:
```
30
4
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Go menyederhanakan perulangan dengan hanya menyediakan satu kata kunci: `for`. Baik perulangan counter klasik, gaya `while`, maupun `infinite loop` semuanya menggunakan `for`, dikombinasikan dengan `break` dan `continue` untuk mengontrol alurnya.
