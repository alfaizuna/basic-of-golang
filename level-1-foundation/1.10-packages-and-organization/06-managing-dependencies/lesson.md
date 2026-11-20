# Managing Dependencies

Mengelola dependensi eksternal menggunakan `go get`, pembersihan otomatis dengan `go mod tidy`, serta verifikasi integritas dengan `go.sum`.

## Contoh

```go
//go:build ignore

package main

import "fmt"

func main() {
	fmt.Println("Perintah CLI dependensi Go:")
	fmt.Println("  go get github.com/google/uuid")
	fmt.Println("  go mod tidy")
}

```

## Penjelasan

1. **`go get <pkg>`**: Mengunduh dan mencatat dependensi ke `go.mod`.
2. **`go mod tidy`**: Menyelaraskan `go.mod` dengan kode sumber yang sebenarnya: menambahkan library yang dipakai dan menghapus library yang tidak lagi digunakan.
3. **`go.sum`**: Berisi hash cryptographic checksum dari setiap modul untuk menjamin bahwa dependensi yang diunduh tidak pernah dimanipulasi atau diubah secara diam-diam (*tamper-proof*).
4. **Semantic Versioning**: Versi modul mengikuti format `vMAJOR.MINOR.PATCH`.

## >_ Exercise

Buat fungsi `FormatSemver(major, minor, patch int, prerelease string) string` yang:
- Jika `prerelease` kosong: kembalikan `"v<major>.<minor>.<patch>"`
- Jika `prerelease` tidak kosong: kembalikan `"v<major>.<minor>.<patch>-<prerelease>"`

### Expected output

```
Release: v1.2.3
Pre-release: v2.0.0-beta.1
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `go mod tidy` menjaga kebersihan file `go.mod` dan `go.sum`.
- `go.sum` menjamin keamanan rantai pasokan kode (*supply-chain security*) melalui checksum hash.
