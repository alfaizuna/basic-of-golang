# Zone Storage

Pada pelajaran kedua, kita akan membuat struktur penyimpanan bernama `Zone`. Sebuah Zone DNS bertanggung jawab menampung kumpulan record yang dimiliki suatu domain.
Karena satu nama domain bisa memiliki beberapa record (misalnya beberapa IP address untuk load balancing, atau kombinasi record yang berbeda), kita menyimpannya dalam struktur `map[string][]*Record`.

## Contoh
example.go
```go
package main

type Zone struct {
	Name    string
	records map[string][]*Record
}

func NewZone(name string) *Zone {
	return &Zone{
		Name:    name,
		records: make(map[string][]*Record),
	}
}

func (z *Zone) RecordCount() int {
	if z == nil || z.records == nil {
		return 0
	}
	total := 0
	for _, recs := range z.records {
		total += len(recs)
	}
	return total
}
```

## Penjelasan
- `records map[string][]*Record`:
  - Kunci map adalah nama domain (misal `"example.com"`).
  - Nilai map adalah slice of pointers `[]*Record`, karena satu domain bisa memiliki lebih dari satu record.
- `NewZone(name string) *Zone`: konstruktor yang mengembalikan pointer ke `Zone` baru dengan map internal `records` yang sudah siap digunakan.
- `(z *Zone) RecordCount() int`: method penghitung total record di seluruh zone dengan nil guard.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.4-project-dns-resolver/02-zone-storage/main.go).

1. Definisikan struct `Zone`:
   - `Name` (`string`)
   - `records` (`map[string][]*Record`)
2. Buat fungsi konstruktor `NewZone(name string) *Zone`:
   - Inisialisasi map internal `records: make(map[string][]*Record)`.
3. Buat method `(z *Zone) RecordCount() int`:
   - Jika `z == nil` atau `z.records == nil`, kembalikan `0`.
   - Iterasi map `z.records` dan jumlahkan panjang slice masing-masing domain (`len(records)`).
   - Kembalikan total jumlah record.

Expected output ketika `go run main.go`:
```
Zone created: example.com (records: 0)
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Menggunakan slice pointer di dalam map (`map[string][]*Record`) memberikan fleksibilitas penuh untuk menyimpan banyak record per nama domain dan melakukan modifikasi data langsung lewat pointer tanpa alokasi ulang.
