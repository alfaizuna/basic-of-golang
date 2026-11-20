# Following CNAME Chains

Record `CNAME` berfungsi sebagai alias untuk nama domain lain. Contohnya:
- `app.example.com` CNAME ➡️ `lb.example.com`
- `lb.example.com` CNAME ➡️ `cluster.example.com`
- `cluster.example.com` A ➡️ `93.184.216.34`

Pada pelajaran keenam, kita membuat method `Resolve(name string) (*Record, error)` yang menelusuri rantai CNAME secara iteratif sampai menemukan record `A` yang menyimpan IP address tujuan.

## Contoh
example.go
```go
package main

import (
	"errors"
)

func (z *Zone) Resolve(name string) (*Record, error) {
	current := name
	for {
		if rec := z.Lookup(current, TypeA); rec != nil {
			return rec, nil
		}

		cname := z.Lookup(current, TypeCNAME)
		if cname == nil {
			return nil, errors.New("record not found")
		}

		current = cname.Value
	}
}
```

## Penjelasan
- Alur resolusi bertingkat:
  1. Cek apakah ada record `TypeA` untuk domain `current`. Jika ada, langsung kembalikan record tersebut.
  2. Jika tidak ada, cek apakah ada record `TypeCNAME`.
  3. Jika ada CNAME, pindahkan `current = cname.Value` dan ulangi iterasi berikutnya.
  4. Jika tidak ada record `A` maupun `CNAME`, kembalikan `errors.New("record not found")`.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.4-project-dns-resolver/06-following-cname-chains/main.go).

Implementasikan method `(z *Zone) Resolve(name string) (*Record, error)`:
- Cari record `TypeA` untuk nama domain saat ini (`current`). Jika ada, kembalikan `rec, nil`.
- Jika tidak ada, cari record `TypeCNAME`.
- Jika ditemukan `CNAME`, ganti domain target dengan nilai `cname.Value` dan ulangi loop.
- Jika tidak ada record yang cocok sama sekali, kembalikan `nil, errors.New("record not found")`.

Expected output ketika `go run main.go`:
```
Resolved app.example.com -> 93.184.216.34
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Penelusuran rantai CNAME memungkinkan resolver menemukan IP target di balik alias berlapis dengan mengikuti pointer dan value antar record.
