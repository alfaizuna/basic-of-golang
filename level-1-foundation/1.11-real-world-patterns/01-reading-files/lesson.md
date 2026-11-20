# Reading Files

Membaca berkas dari filesystem secara aman menggunakan `os.ReadFile` dan menangani error jika file tidak ditemukan.

## Contoh

```go
//go:build ignore

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("example.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File tidak ditemukan!")
		return
	} else if err != nil {
		fmt.Println("Gagal membaca file:", err)
		return
	}
	fmt.Println("Isi file:", string(data))
}

```

## Penjelasan

1. **`os.ReadFile(path)`**: Cara tercepat membaca seluruh isi berkas ke dalam memory (`[]byte`). Fungsi ini secara otomatis membuka dan menutup file.
2. **`errors.Is(err, os.ErrNotExist)`**: Cara modern dan idiomatis untuk mengecek secara spesifik apakah kegagalan disebabkan karena berkas tidak ada di disk.
3. **Konversi ke String**: Karena hasilnya bertipe `[]byte`, ubah menjadi string dengan `string(data)` jika berisi teks.

## >_ Exercise

Buat fungsi `LoadConfigText(path string) (string, error)` yang:
1. Membaca berkas pada `path` dengan `os.ReadFile`.
2. Jika file tidak ada (`errors.Is(err, os.ErrNotExist)`), kembalikan `"", fmt.Errorf("file not found")`.
3. Jika ada error lain, kembalikan `"", err`.
4. Jika sukses, kembalikan isi file sebagai string dan `nil`.

### Expected output

```
Loaded config: OK
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `os.ReadFile` membaca seluruh file ke dalam `[]byte` dalam satu operasi.
- Gunakan `errors.Is(err, os.ErrNotExist)` untuk mendeteksi file yang hilang.
