# Packages

Package adalah unit modularisasi kode paling mendasar di Go. Semua file dalam satu folder harus mendeklarasikan package yang sama.

## Contoh

```go
//go:build ignore

package main

import "fmt"

func main() {
	fmt.Println("Ini adalah package main yang menghasilkan biner executable")
}

```

## Penjelasan

1. **`package main`**: Menandakan bahwa package ini adalah titik masuk (*entry point*) aplikasi yang dapat dikompilasi menjadi berkas biner eksekusi (*executable*). Wajib memiliki fungsi `func main()`.
2. **Package Library**: Diberi nama sesuai fungsionalitasnya (misalnya `package mathutil` atau `package storage`). Digunakan untuk diimpor oleh file atau package lain dan tidak memiliki `func main()`.
3. **Aturan 1 Folder = 1 Package**: Semua file `.go` yang berada dalam satu direktori yang sama **harus** memiliki nama deklarasi package yang identik.

## >_ Exercise

Buat fungsi `DescribePackage(name string, isMain bool) string` yang mengembalikan:
- `"<name> (executable)"` jika `isMain == true`
- `"<name> (library)"` jika `isMain == false`

### Expected output

```
main (executable)
mathutil (library)
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `package main` menghasilkan program eksekusi; nama package lain menjadi modul library.
- Semua file dalam folder yang sama wajib memakai nama package yang sama.
