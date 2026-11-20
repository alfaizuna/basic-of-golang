# Hints

### Hint 1: Arah Umum
Di Go, argumen parameter fungsi adalah variabel lokal di dalam fungsi tersebut, sehingga kamu bisa langsung mengambil alamatnya menggunakan operator `&`.
Untuk menerapkan default, periksa apakah pointer field bernilai `nil`. Jika tidak `nil`, ambil nilainya menggunakan operator dereference `*`.

### Hint 2: Sintaks yang Relevan
- Di `StringPtr(s string) *string`: kamu bisa langsung mengembalikan `&s`.
- Di `ResolveServerConfig`: gunakan `if cfg.Host != nil { res.Host = *cfg.Host }` untuk masing-masing field (`Host`, `Port`, `Debug`).

### Hint 3: Kerangka Kode
```go
func StringPtr(s string) *string {
    return &s
}

// Di ResolveServerConfig:
if cfg.Host != nil {
    res.Host = *cfg.Host
}
if cfg.Port != nil {
    res.Port = *cfg.Port
}
if cfg.Debug != nil {
    res.Debug = *cfg.Debug
}
```
