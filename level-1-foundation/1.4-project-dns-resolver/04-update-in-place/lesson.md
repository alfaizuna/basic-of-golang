# Update in Place

Pada pelajaran keempat, kita memanfaatkan keunggulan pointer untuk memperbarui data record secara langsung di tempat (*in-place mutation*), tanpa perlu menghapus atau menulis ulang isi map.

## Contoh
example.go
```go
package main

func (r *Record) UpdateValue(newValue string) {
	if r == nil {
		return
	}
	r.Value = newValue
}

func (r *Record) UpdateTTL(newTTL int) {
	if r == nil {
		return
	}
	r.TTL = newTTL
}
```

## Penjelasan
- Pointer receiver `(r *Record)` memungkinkan kita memodifikasi field struct `Record` yang tersimpan di dalam memori zone.
- Siapapun yang memegang pointer atau memanggil `Lookup` berikutnya akan langsung melihat nilai baru tersebut secara otomatis.
- Method `UpdateRecord(name string, recordType RecordType, newValue string, newTTL int) bool` pada `Zone`:
  - Menggunakan `Lookup` internal untuk mendapatkan pointer record.
  - Memperbarui value dan TTL jika ditemukan.
  - Mengembalikan status `true` jika berhasil, atau `false` jika record tidak ditemukan.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.4-project-dns-resolver/04-update-in-place/main.go).

Lengkapi method berikut:
1. Pada `Record`:
   - `(r *Record) UpdateValue(newValue string)`
   - `(r *Record) UpdateTTL(newTTL int)`
2. Pada `Zone`:
   - `(z *Zone) UpdateRecord(name string, recordType RecordType, newValue string, newTTL int) bool`:
     - Cari record dengan `z.Lookup(name, recordType)`.
     - Jika tidak ditemukan (`rec == nil`), kembalikan `false`.
     - Perbarui value dan TTL record menggunakan method di atas (atau assign langsung ke field).
     - Kembalikan `true`.

Expected output ketika `go run main.go`:
```
Before: example.com -> 93.184.216.34 (TTL: 300)
Update success: true
After : example.com -> 93.184.216.35 (TTL: 600)
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Mengubah state melalui pointer receiver menghilangkan kebutuhan untuk menghapus dan menambahkan ulang elemen di dalam map atau slice.
