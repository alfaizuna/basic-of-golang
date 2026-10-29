# Strings Under the Hood

Here's how Go stores strings as immutable byte slices and handles UTF-8 text.

## Contoh
```go
// example.go
//go:build ignore

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	greeting := "Halo 世界"

	// len() menghitung BYTE, bukan karakter!
	fmt.Println("Jumlah byte:", len(greeting)) // 11 byte

	// utf8.RuneCountInString() menghitung jumlah karakter sebenarnya
	fmt.Println("Jumlah rune:", utf8.RuneCountInString(greeting)) // 7 rune

	// Paket strings menyediakan berbagai fungsi utilitas teks
	fmt.Println(strings.ToLower("GO LANGUAGE"))
	fmt.Println(strings.TrimSpace("  bersih  "))
}
```

## Penjelasan
- String di Go pada dasarnya adalah *read-only slice of bytes* ber-encoding UTF-8 dan bersifat *immutable* (tidak dapat diubah per karakternya langsung).
- Fungsi bawaan `len(s)` mengembalikan **jumlah byte**, bukan jumlah karakter jika string mengandung karakter multi-byte seperti emoji atau aksara non-Latin.
- Untuk menghitung jumlah karakter (*rune*) sesungguhnya, gunakan fungsi `utf8.RuneCountInString(s)` dari paket standard `"unicode/utf8"`.
- Perulangan `for idx, r := range s` otomatis membaca string karakter per karakter (*rune*) dan mengembalikan indeks byte awal dari karakter tersebut.
- Paket standard `"strings"` menyediakan banyak fungsi pengolahan teks penting seperti `strings.TrimSpace`, `strings.ToLower`, `strings.ToUpper`, `strings.ReplaceAll`, dan `strings.Contains`.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.1-go-basics/15-strings-under-the-hood/main.go):
1. Selesaikan fungsi `countBytesAndRunes(text string) (int, int)`:
   - Kembalikan `len(text)` (jumlah byte) dan `utf8.RuneCountInString(text)` (jumlah rune).
2. Selesaikan fungsi `normalizeText(text string) string`:
   - Bersihkan spasi di awal dan akhir string menggunakan `strings.TrimSpace(text)`.
   - Ubah seluruh huruf menjadi huruf kecil menggunakan `strings.ToLower(...)`.
   - Kembalikan teks yang sudah dinormalisasi.
3. Di dalam `func main()`:
   - Panggilan pertama `countBytesAndRunes("Go 🚀")` sudah disediakan dan akan mencetak `7 4` setelah kamu menyelesaikan fungsi tersebut.
   - Panggil `normalizeText("  Belajar GOLANG  ")`, simpan ke variabel `cleaned`, lalu cetak dengan `fmt.Println(cleaned)`.

Expected output:
```
7 4
belajar golang
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
String di Go adalah *immutable slice of bytes* ber-encoding UTF-8. Selalu ingat bahwa `len(s)` mengukur ukuran memori dalam byte, sedangkan `utf8.RuneCountInString(s)` mengukur banyaknya karakter Unicode.
