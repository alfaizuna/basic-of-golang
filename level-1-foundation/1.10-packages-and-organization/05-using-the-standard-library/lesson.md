# Using the Standard Library

Standard library Go terkenal sangat lengkap, stabil, dan berkinerja tinggi, menjadikannya andalan utama pengembang Go (*batteries included*).

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	s := strings.TrimSpace("   42   ")
	n, _ := strconv.Atoi(s)
	fmt.Printf("Angka: %d, Tanggal: %s\n", n, time.Now().Format("2006-01-02"))
}

```

## Penjelasan

1. **Kekuatan Standard Library**: Go memiliki pustaka standar kelas dunia untuk manipulasi string (`strings`), konversi data (`strconv`), kriptografi (`crypto`), HTTP server/client (`net/http`), JSON (`encoding/json`), dan I/O (`io`, `os`).
2. **Dokumentasi `pkg.go.dev`**: Dokumentasi resmi seluruh pustaka standar dapat dibaca daring di https://pkg.go.dev atau via terminal dengan perintah `go doc <package>`.

## >_ Exercise

Buat fungsi `TransformInput(raw string) (int, string, error)` yang:
1. Memotong spasi di awal dan akhir `raw` menggunakan `strings.TrimSpace`.
2. Mengonversi teks bersih tersebut ke integer menggunakan `strconv.Atoi`. Jika ada error, kembalikan `0, "", err`.
3. Mengembalikan angka tersebut, string dalam huruf besar semua (`strings.ToUpper(raw)`), dan `nil`.

### Expected output

```
Num: 250, Upper: 250 USD
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Prioritaskan penggunaan standard library Go sebelum menambah third-party dependencies.
- `strings` dan `strconv` menyediakan fungsi utilitas teks yang aman dan efisien.
