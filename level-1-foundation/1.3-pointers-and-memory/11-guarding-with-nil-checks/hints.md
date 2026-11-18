# Hints

### Hint 1: Arah Umum
Untuk mengamankan pointer dari panic, gunakan *guard clause* dengan `if ptr == nil`.
Untuk fungsi pencarian yang berpotensi tidak menemukan hasil, gunakan `errors.New(...)` untuk membuat pesan error baru dan kembalikan `nil` untuk pointernya.

### Hint 2: Sintaks yang Relevan
- Di `FormatUserName`: cek apakah `u == nil` di awal fungsi, jika ya langsung `return "Guest"`.
- Di `FindUser`: jika loop selesai tanpa menemukan ID yang cocok, kembalikan `nil, errors.New("user not found")`.

### Hint 3: Kerangka Kode
```go
// Di FormatUserName:
if u == nil {
    return "Guest"
}

// Di FindUser (bagian akhir):
return nil, errors.New("user not found")
```
