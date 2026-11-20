# Course 1.10: Packages & Organization

## Tujuan
Memahami organisasi kode profesional di Go: perbedaan package executable (`package main`) dan library, satu folder = satu package, inisialisasi dan manajemen Go Modules (`go.mod` dan `go.sum`), enkapsulasi melalui identifier exported (kapital) vs unexported (huruf kecil), teknik impor package lokal dan alias import, eksplorasi pustaka standar Go, manajemen dependensi pihak ketiga (`go get`, `go mod tidy`), serta struktur arsitektur proyek Go skala besar (`cmd/`, `internal/`) dan pencegahan import cycle.

## Prasyarat
Telah menyelesaikan Course 1.1 s.d. Course 1.9.

## Daftar Lesson

| # | Lesson | Fokus | Status |
|---|--------|-------|--------|
| 01 | Packages | `package main` vs package library, satu folder = satu package | 🔒 |
| 02 | Go Modules | `go mod init`, isi `go.mod`, path module | 🔒 |
| 03 | Exported vs Unexported | huruf kapital = publik, enkapsulasi dan API kecil | 🔒 |
| 04 | Importing Your Own Packages | multi-package dalam module, import alias | 🔒 |
| 05 | Using the Standard Library | memanfaatkan library bawaan (`strconv`, `strings`, `math`) | 🔒 |
| 06 | Managing Dependencies | semver, `go get`, `go mod tidy`, `go.sum` | 🔒 |
| 07 | Structuring Larger Projects | layout standar `cmd/` dan `internal/`, pencegahan circular import | 🔒 |

## Hasil Akhir
Kemampuan menstrukturkan codebase Go dari skala skrip tunggal hingga proyek berskala besar secara teratur, modular, dan mematuhi konvensi komunitas Go.
