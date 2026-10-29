# Format Like a Pro

Here's how to format strings, numbers, and types in Go like a pro.

## Contoh
```go
// example.go
//go:build ignore

package main

import "fmt"

func main() {
	nama := "Budi"
	umur := 25
	saldo := 150000.75

	// %T untuk tipe, %q untuk kutip otomatis, %.2f untuk presisi angka desimal
	fmt.Printf("Tipe: %T | Nama: %q | Saldo: Rp %.2f\n", saldo, nama, saldo)

	// fmt.Sprintf menyimpan hasil format ke dalam variabel string
	pesan := fmt.Sprintf("User: %s (Umur: %d)", nama, umur)
	fmt.Println(pesan)
}
```

## Penjelasan
- `fmt.Printf` mencetak langsung ke terminal (*stdout*), sedangkan `fmt.Sprintf` merangkai dan **mengembalikan** hasilnya sebagai nilai `string`.
- `%v` → format nilai default untuk sembarang tipe data.
- `%T` → mencetak nama tipe data dari variabel (misal `int`, `string`, `float64`).
- `%s` → mencetak string biasa; `%q` → mencetak string berbalut tanda kutip ganda (`"..."`), sangat berguna untuk melihat teks yang mengandung spasi atau karakter tak kasat mata.
- `%d` → integer basis 10; `%t` → boolean (`true` / `false`).
- `%.2f` → angka pecahan (*float*) dibatasi 2 digit di belakang koma.
- `%04d` → integer dengan lebar minimal 4 digit dan diisi nol di depan (*zero-padding*, misal `7` menjadi `0007`).

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.1-go-basics/16-format-like-a-pro/main.go):
1. Selesaikan fungsi `formatProduct(name string, price float64, stock int) string`:
   - Gunakan `fmt.Sprintf` untuk menghasilkan format string:
     `Product: %q | Price: $%.2f | Stock: %d`
2. Selesaikan fungsi `describeValue(val int) string`:
   - Gunakan `fmt.Sprintf` untuk menghasilkan format string:
     `Type: %T | Padded: %04d`
   - Berikan argumen `val` untuk kedua placeholder (`%T` dan `%04d`).
3. Di dalam `func main()`:
   - Panggilan pertama `formatProduct("Mechanical Keyboard", 89.9, 15)` sudah disediakan dan akan mencetak hasil formatnya setelah fungsi diselesaikan.
   - Panggil `describeValue(42)`, simpan ke `d`, lalu cetak dengan `fmt.Println(d)`.

Expected output:
```
Product: "Mechanical Keyboard" | Price: $89.90 | Stock: 15
Type: int | Padded: 0042
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Gunakan `fmt.Sprintf` saat ingin merangkai string terformat ke dalam variabel. Kuasai verb penting seperti `%v`, `%T`, `%q`, serta pengatur presisi `%.2f` dan padding `%04d` untuk menghasilkan teks dan log yang rapi.
