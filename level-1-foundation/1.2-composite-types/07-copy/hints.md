### Hint 1
Gunakan fungsi bawaan `make([]string, len(original))` untuk menyiapkan slice `backup` agar panjangnya sesuai sebelum memanggil `copy(backup, original)`.

### Hint 2
Sintaks membuat dan menyalin:
```go
backup := make([]string, len(original))
copy(backup, original)
```
Untuk mengubah elemen terakhir, akses indeks terakhir (`len(backup) - 1` atau indeks `2`).

### Hint 3
Kerangka kode lengkap:
```go
backup := make([]string, len(original))
copy(backup, original)
backup[2] = "TypeScript"

fmt.Println("Original:", original)
fmt.Println("Backup:", backup)
```
