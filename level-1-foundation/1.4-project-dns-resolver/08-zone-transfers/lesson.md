# Zone Transfers

Pada pelajaran terakhir Course 1.4 ini, kita akan mengimplementasikan **Zone Transfer** (AXFR dalam protokol DNS), yaitu proses menyalin seluruh isi zone dari server DNS utama (*primary*) ke server sekunder (*replica/backup*).

Di sinilah pemahaman memori sangat diuji:
- **Shallow Copy (Berbahaya)**: jika kita hanya menyalin pointer `*Record` ke zone baru, maka kedua zone akan menunjuk ke alamat memori struct yang sama. Akibatnya, pembaruan di zone replika akan ikut mengubah data di zone utama!
- **Deep Copy (Benar & Aman)**: kita harus mengalokasikan struct `Record` baru untuk setiap record yang ditransfer, sehingga kedua zone sepenuhnya independen di memori.

## Contoh
example.go
```go
package main

// Deep Copy: Alokasikan struct Record baru untuk setiap entri
func (z *Zone) Transfer() *Zone {
	if z == nil {
		return nil
	}
	target := NewZone(z.Name)
	for _, recs := range z.records {
		for _, r := range recs {
			copiedRec := &Record{
				Name:  r.Name,
				Type:  r.Type,
				Value: r.Value,
				TTL:   r.TTL,
			}
			target.AddRecord(copiedRec)
		}
	}
	return target
}
```

## Penjelasan
- `copiedRec := &Record{...}`: mengalokasikan struct baru di heap dengan menyalin nilai field-fieldnya (*value copy*).
- `target.AddRecord(copiedRec)`: menyimpan pointer baru ke dalam zone sekunder.
- Alamat memori `copiedRec` berbeda dari `r` (`copiedRec != r`), sehingga kedua zone tidak berbagi pointer.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.4-project-dns-resolver/08-zone-transfers/main.go).

Implementasikan method `(z *Zone) Transfer() *Zone`:
1. Jika `z == nil`, kembalikan `nil`.
2. Buat zone baru dengan `target := NewZone(z.Name)`.
3. Lakukan perulangan bersarang pada `z.records`:
   - Buat salinan struct `Record` baru dengan nilai `Name`, `Type`, `Value`, dan `TTL` yang sama.
   - Tambahkan salinan record tersebut ke `target.AddRecord(...)`.
4. Kembalikan pointer `target` zone baru.

Expected output ketika `go run main.go`:
```
Primary zone count: 2
Replica zone count: 2
Independent memory confirmed: true
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Selamat! Kamu telah menuntaskan seluruh materi Course 1.4: Project DNS Resolver!
Kamu kini menguasai pemodelan record dengan pointer, zone storage map-of-slice, mutasi in-place, caching dengan mock clock, resolusi rekursif rantai CNAME dengan cycle detection, serta pentingnya deep copy pada transfer data.
