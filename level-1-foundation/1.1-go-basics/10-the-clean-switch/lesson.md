# The Clean Switch

Here's your first look at Go's cleaner switch statement, which needs no break statements and can run without an expression.

## Contoh
```go
// example.go
package main

import "fmt"

func main() {
	role := "admin"

	// 1. Switch biasa (dengan ekspresi nilai)
	switch role {
	case "admin":
		fmt.Println("Akses penuh")
	case "editor", "author": // multiple values dalam satu case
		fmt.Println("Akses edit")
	default:
		fmt.Println("Akses tamu")
	}

	score := 85

	// 2. Switch tanpa ekspresi (seperti rantai if-else yang rapi)
	switch {
	case score >= 90:
		fmt.Println("Grade A")
	case score >= 80:
		fmt.Println("Grade B")
	default:
		fmt.Println("Grade C")
	}
}
```

## Penjelasan
- `switch role { ... }` → Mencocokkan nilai variabel `role` dengan nilai di tiap `case`.
- Tidak perlu `break` → Go secara otomatis berhenti di akhir `case` yang cocok (tidak ada *fallthrough* otomatis).
- `case "editor", "author":` → Satu `case` dapat memeriksa beberapa nilai sekaligus dipisahkan tanda koma.
- `switch { ... }` → Switch tanpa ekspresi mengevaluasi kondisi boolean di tiap `case`, menjadikannya alternatif yang lebih rapi dibanding rantai `if-else if-else` yang panjang.
- `default:` → Bagian opsional yang dieksekusi jika tidak ada satupun `case` yang cocok.

## >_ Exercise
Buka file `main.go`:
1. Buat fungsi `dayType(day string) string`:
   - Gunakan `switch day` (switch dengan ekspresi):
     - Untuk `"senin"`, `"selasa"`, `"rabu"`, `"kamis"`, `"jumat"`, kembalikan `"Hari Kerja"`.
     - Untuk `"sabtu"`, `"minggu"`, kembalikan `"Akhir Pekan"`.
     - Selain itu (`default`), kembalikan `"Tidak Dikenal"`.
2. Buat fungsi `categorizeAge(age int) string`:
   - Gunakan `switch` tanpa ekspresi (`switch { ... }`):
     - Jika `age < 13`, kembalikan `"Anak-anak"`.
     - Jika `age <= 19`, kembalikan `"Remaja"`.
     - Selain itu (`default`), kembalikan `"Dewasa"`.
3. Di dalam `func main()`:
   - Panggil `dayType("sabtu")` dan cetak hasilnya dengan `fmt.Println`.
   - Panggil `categorizeAge(17)` dan cetak hasilnya dengan `fmt.Println`.

Expected output:
```
Akhir Pekan
Remaja
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Di Go, `switch` tidak memerlukan `break` di akhir setiap `case`. Satu `case` bisa menampung banyak nilai yang dipisahkan koma, dan `switch` tanpa ekspresi (`switch {}`) menyediakan cara yang bersih dan elegan untuk menggantikan rantai `if-else` bertingkat.
