# Writing Files

Menulis data ke filesystem menggunakan `os.WriteFile` dan `os.Create` dengan idiom penutupan `defer f.Close()`.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"os"
)

func main() {
	// Menulis data langsung
	err := os.WriteFile("output.txt", []byte("Halo dari Go!"), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File berhasil ditulis")
}

```

## Penjelasan

1. **`os.WriteFile(path, data, perm)`**: Membuat atau menimpa (*overwrite*) seluruh berkas dengan slice byte. `0644` adalah izin file standar di Unix (pemilik bisa baca/tulis, yang lain hanya baca).
2. **`os.Create(path)`**: Membuat berkas baru dan mengembalikan pointer berkas `*os.File`.
3. **Idiom `defer f.Close()`**: Selalu tempatkan `defer f.Close()` tepat setelah memeriksa error `os.Create` agar file descriptor tidak bocor (*resource leak*).

## >_ Exercise

Buat fungsi `SaveReport(filename string, content string) error` yang:
Menulis `content` ke `filename` dengan izin `0644` menggunakan `os.WriteFile` dan mengembalikan error jika ada.

### Expected output

```
Report saved successfully
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `os.WriteFile` membuat atau menimpa berkas secara atomik untuk operasi sederhana.
- Untuk operasi bertahap, gunakan `os.Create` disertai `defer f.Close()`.
