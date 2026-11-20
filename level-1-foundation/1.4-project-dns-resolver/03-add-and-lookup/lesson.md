# Add & Lookup

Pada pelajaran ketiga, kita akan menambahkan operasi fundamental ke dalam `Zone`: menambahkan record (`AddRecord`) dan mencari record (`Lookup`).

## Contoh
example.go
```go
package main

func (z *Zone) AddRecord(rec *Record) {
	if z == nil || rec == nil {
		return
	}
	z.records[rec.Name] = append(z.records[rec.Name], rec)
}

func (z *Zone) Lookup(name string, recordType RecordType) *Record {
	if z == nil || z.records == nil {
		return nil
	}
	for _, r := range z.records[name] {
		if r.Type == recordType {
			return r
		}
	}
	return nil
}
```

## Penjelasan
- `AddRecord(rec *Record)`:
  - Menyimpan pointer `rec` ke slice `z.records[rec.Name]` menggunakan `append`.
- `Lookup(name string, recordType RecordType) *Record`:
  - Mengambil slice record untuk domain `name`.
  - Mencari record dengan tipe yang cocok (`r.Type == recordType`).
  - Mengembalikan pointer `*Record` jika ditemukan, atau `nil` jika tidak ditemukan (*nil safe*).

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.4-project-dns-resolver/03-add-and-lookup/main.go).

Lengkapi dua method pada `Zone`:
1. `(z *Zone) AddRecord(rec *Record)`:
   - Tambahkan *nil guard* (`if z == nil || rec == nil { return }`).
   - Tambahkan `rec` ke dalam `z.records[rec.Name]` menggunakan `append`.
2. `(z *Zone) Lookup(name string, recordType RecordType) *Record`:
   - Tambahkan *nil guard* (`if z == nil || z.records == nil { return nil }`).
   - Cari di slice `z.records[name]`. Jika ditemukan record dengan tipe yang cocok, kembalikan pointernya.
   - Jika tidak ada yang cocok, kembalikan `nil`.

Expected output ketika `go run main.go`:
```
Lookup example.com (A): 93.184.216.34
Lookup notfound.com (A): nil
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Pencarian yang mengembalikan pointer (`*Record` atau `nil`) mempermudah pengecekan keberadaan data (*nil check*) sekaligus menyediakan akses langsung ke record asli untuk pembaruan berikutnya.
