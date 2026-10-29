# Slug Generator 3: Transform

Here's the third step in building your CLI Slug Generator: transforming cleaned text into a URL-friendly slug by converting separators to dashes and collapsing duplicates.

## Contoh
```go
// example.go
//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func collapseSeparators(input string) string {
	var result string
	inSpace := false

	for _, r := range input {
		if r == ' ' {
			if !inSpace && len(result) > 0 {
				result += "-"
				inSpace = true
			}
		} else {
			result += string(r)
			inSpace = false
		}
	}
	return strings.Trim(result, "-")
}

func main() {
	sample := "  belajar   golang   dasar  "
	fmt.Println("Result:", collapseSeparators(sample)) // belajar-golang-dasar
}
```

## Penjelasan
- Slug URL harus memisahkan setiap kata dengan tanda strip (`-`), tanpa spasi atau strip ganda seperti `---`.
- Teknik *flag boolean* (misalnya `inSeparator := false`) berguna untuk mendeteksi apakah kita sedang berada di dalam urutan pemisah, sehingga spasi/strip beruntun bisa diciutkan (*collapse*) menjadi satu strip saja.
- Kondisi `len(result) > 0` mencegah strip ditambahkan di paling awal string.
- Fungsi `strings.Trim(result, "-")` memotong sisa strip yang mungkin tertinggal di ujung awal atau akhir teks.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.1-go-basics/slug-generator/main.go):
1. Tambahkan fungsi baru `transformToSlug(text string) string`:
   - Mulai dengan memanggil `cleaned := cleanText(text)` untuk memperoleh teks yang sudah dibersihkan.
   - Siapkan variabel penampung `var result string` dan *flag* `inSeparator := false`.
   - Iterasi setiap karakter `r` pada `cleaned`:
     - Jika `r == ' ' || r == '-'`:
       - Jika `!inSeparator && len(result) > 0`, tambahkan `"-"` ke `result`, dan ubah `inSeparator = true`.
     - Jika karakter alfanumerik biasa:
       - Tambahkan karakter ke `result` (`result += string(r)`), dan set `inSeparator = false`.
   - Kembalikan `strings.Trim(result, "-")`.
2. Di dalam `func main()`:
   - Ganti isi fungsi `main()` dengan:
     ```go
     input := "  Go 101: Belajar #1!  "
     slug := transformToSlug(input)
     fmt.Printf("Input: %q\n", input)
     fmt.Printf("Slug:  %q\n", slug)
     ```

Expected output:
```
Input: "  Go 101: Belajar #1!  "
Slug:  "go-101-belajar-1"
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Transformasi slug menuntut penanganan karakter pemisah yang cermat. Menggunakan variabel status (*state flag*) saat iterasi memungkinkan kita menyatukan beberapa pemisah beruntun menjadi satu pemisah yang rapi.
