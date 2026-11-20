# Go Modules

Go Modules adalah sistem manajemen dependensi resmi di Go yang mendefinisikan root direktori, versi Go, dan seluruh library eksternal.

## Contoh

```go
//go:build ignore

package main

import "fmt"

func main() {
	fmt.Println("go.mod mendefinisikan modul: module github.com/user/myproject")
}

```

## Penjelasan

1. **`go mod init <path>`**: Perintah untuk membuat file `go.mod` di root proyek.
2. **Komponen `go.mod`**:
   - `module <path>`: identitas unik modul (biasanya path repositori seperti `github.com/org/repo`).
   - `go <version>`: batas minimum versi Go compiler.
   - `require (...)`: daftar dependensi pihak ketiga beserta versinya.
3. **Module Root**: Semua impor internal dimulai dari path module yang didefinisikan di baris pertama `go.mod`.

## >_ Exercise

Buat fungsi `FormatModuleHeader(modulePath, goVersion string) string` yang menghasilkan format string 2 baris persis seperti permulaan file `go.mod`:
```
module <modulePath>

go <goVersion>
```

### Expected output

```
module github.com/user/myapp

go 1.23
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `go.mod` adalah fondasi dependensi Go yang menandai akar dari pohon modul.
- Path modul menjadi awalan untuk semua subpackage internal.
