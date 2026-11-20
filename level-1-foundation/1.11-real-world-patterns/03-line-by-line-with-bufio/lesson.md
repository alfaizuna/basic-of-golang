# Line-by-Line with bufio

Membaca aliran data atau berkas besar baris demi baris secara efisien tanpa memuat seluruh berkas ke memori menggunakan `bufio.Scanner`.

## Contoh

```go
//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	raw := "Baris 1\nBaris 2\nBaris 3"
	scanner := bufio.NewScanner(strings.NewReader(raw))

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Dibaca:", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error saat scanning:", err)
	}
}

```

## Penjelasan

1. **Kenapa `bufio.Scanner`**: Jika berkas berukuran ratusan Megabyte atau Gigabyte, `os.ReadFile` akan menghabiskan RAM. `bufio.Scanner` hanya membaca satu baris kecil ke buffer pada satu waktu.
2. **Loop `scanner.Scan()`**: Mengembalikan `true` selama masih ada baris berikutnya, dan `false` jika mencapai akhir file (EOF) atau error.
3. **`scanner.Text()`**: Mengambil baris saat ini dalam bentuk string tanpa karakter baris baru (`
`).
4. **`scanner.Err()`**: Wajib dicek setelah loop untuk memastikan loop berhenti karena EOF normal, bukan karena kegagalan I/O.

## >_ Exercise

Buat fungsi `FilterLines(r io.Reader, keyword string) ([]string, error)` yang:
1. Membaca baris demi baris dari `r` menggunakan `bufio.NewScanner`.
2. Menyimpan baris-baris yang mengandung `keyword` (`strings.Contains(line, keyword)`).
3. Setelah scan selesai, periksa `scanner.Err()`. Jika ada error, kembalikan `nil, err`.
4. Jika sukses, kembalikan slice baris yang cocok.

### Expected output

```
Matched lines count: 2
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `bufio.Scanner` adalah alat standar industri untuk membaca stream teks berukuran besar tanpa memory overhead.
- Selalu periksa `scanner.Err()` setelah loop selesai.
