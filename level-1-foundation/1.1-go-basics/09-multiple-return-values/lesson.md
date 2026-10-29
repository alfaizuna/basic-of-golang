# Multiple Return Values

Here's your first look at returning multiple values from a function and ignoring unused values with the blank identifier.

## Contoh
```go
// example.go
package main

import "fmt"

// Fungsi yang mengembalikan 2 nilai: hasil bagi dan sisa bagi
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	// Menampung kedua nilai kembalian
	q, r := divide(10, 3)
	fmt.Printf("Hasil: %d, Sisa: %d\n", q, r)

	// Mengabaikan sisa bagi menggunakan blank identifier (_)
	onlyQuotient, _ := divide(10, 3)
	fmt.Println("Hanya hasil:", onlyQuotient)
}
```

## Penjelasan
- `(int, int)` → Tipe return lebih dari satu wajib dibungkus dalam tanda kurung `(...)` setelah daftar parameter.
- `return quotient, remainder` → Nilai-nilai kembalian dipisahkan dengan koma sesuai urutan tipenya.
- `q, r := divide(10, 3)` → Menampung semua nilai kembalian ke beberapa variabel sekaligus.
- `_` (blank identifier) → Karakter garis bawah digunakan untuk mengabaikan nilai kembalian yang tidak kita butuhkan agar tidak menyebabkan error *unused variable*.

## >_ Exercise
Buka file `main.go`:
1. Buat fungsi `calculate(a, b int) (int, int)` di luar `main()`:
   - Mengembalikan 2 nilai: hasil penjumlahan (`a + b`) dan pengurangan (`a - b`).
2. Di dalam `func main()`:
   - Panggil `calculate(20, 5)`, simpan hasilnya ke dua variabel: `sum` dan `diff`.
   - Cetak kedua variabel tersebut: `fmt.Println(sum, diff)`.
   - Panggil `calculate(50, 10)`, simpan hasil pertama (penjumlahan) ke variabel `total`, dan abaikan hasil kedua (pengurangan) menggunakan blank identifier (`_`).
   - Cetak variabel tersebut: `fmt.Println(total)`.

Expected output:
```
25 15
60
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Go mendukung fungsi dengan banyak nilai kembalian (*multiple return values*) dengan menuliskan tipe-tipe kembalian di dalam tanda kurung `(...)`. Gunakan blank identifier `_` untuk mengabaikan nilai return yang tidak kamu perlukan tanpa memicu error variabel yang tidak digunakan.
