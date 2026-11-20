# Structuring Larger Projects

Menerapkan struktur folder standar komunitas Go (`cmd/`, `internal/`, `pkg/`) dan mencegah kesalahan desain dependensi sirkular (import cycle).

## Contoh

```go
//go:build ignore

package main

import "fmt"

func main() {
	fmt.Println("Layout standar proyek Go:")
	fmt.Println("  cmd/          -> entry point biner (misal cmd/api/main.go)")
	fmt.Println("  internal/     -> kode privat yang dilarang diimpor modul lain")
	fmt.Println("  pkg/          -> library publik yang aman dipakai modul luar")
}

```

## Penjelasan

1. **Direktori `cmd/`**: Setiap subdirektori di bawah `cmd/` berisi `main.go` untuk satu biner eksekusi (misalnya `cmd/server/main.go`, `cmd/cli/main.go`).
2. **Direktori `internal/`**: Ditegakkan langsung oleh Go compiler! Paket di dalam `internal/` hanya boleh diimpor oleh kode di dalam repositori/modul yang sama. Proyek eksternal yang mencoba mengimpor akan ditolak kompilator.
3. **Pencegahan Import Cycle**: Go **melarang keras** circular import (`package A import B`, dan `package B import A`). Jika terjadi, pecah dependensi bersama ke package ketiga atau gunakan interface.

## >_ Exercise

Buat fungsi `ValidateImportCycle(dependencies map[string][]string, start string) bool` yang:
Mendeteksi apakah terdapat siklus (cycle) dari package `start` menggunakan penelusuran graf sederhana (DFS).
Kembalikan `true` jika terdeteksi cycle (siklus ketergantungan ilegal), atau `false` jika graf dependensi bersih (Acyclic).

### Expected output

```
Has cycle: true
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `internal/` dilindungi oleh compiler agar tidak bisa diimpor pihak luar.
- Hindari circular dependency dengan interface atau memisahkan modul bersama.
