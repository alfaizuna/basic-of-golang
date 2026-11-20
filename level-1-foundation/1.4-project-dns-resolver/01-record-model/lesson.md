# Record Model

Selamat datang di Course 1.4: Project DNS Resolver! Di course ini, kita akan membangun engine DNS Resolver in-memory.
Pada pelajaran pertama ini, kita akan memodelkan struktur dasar sebuah DNS Record dan konstanta tipe record-nya.

## Contoh
example.go
```go
package main

type RecordType string

const (
	TypeA     RecordType = "A"
	TypeCNAME RecordType = "CNAME"
)

type Record struct {
	Name  string
	Type  RecordType
	Value string
	TTL   int
}

func NewRecord(name string, recordType RecordType, value string, ttl int) *Record {
	return &Record{
		Name:  name,
		Type:  recordType,
		Value: value,
		TTL:   ttl,
	}
}
```

## Penjelasan
- `RecordType string`: custom type berbasis string untuk menyatakan tipe record DNS secara aman (*type safety*).
- `const TypeA RecordType = "A"`: record tipe `A` memetakan nama domain ke alamat IPv4 (contoh: `"example.com"` → `"93.184.216.34"`).
- `const TypeCNAME RecordType = "CNAME"`: record `CNAME` (Canonical Name) memetakan nama domain alias ke domain lain (contoh: `"www.example.com"` → `"example.com"`).
- `Record`: struct penyimpan data DNS record yang berisi `Name`, `Type`, `Value`, dan `TTL` (Time To Live dalam satuan detik).
- `NewRecord(...) *Record`: fungsi konstruktor yang mengembalikan pointer ke `Record` baru di heap.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.4-project-dns-resolver/01-record-model/main.go).

1. Buat tipe data custom `RecordType` berbasis `string`:
   ```go
   type RecordType string
   ```
2. Definisikan dua konstanta:
   - `TypeA` dengan nilai `"A"`
   - `TypeCNAME` dengan nilai `"CNAME"`
3. Definisikan struct `Record` dengan field:
   - `Name` (`string`)
   - `Type` (`RecordType`)
   - `Value` (`string`)
   - `TTL` (`int`)
4. Buat fungsi konstruktor `NewRecord(name string, recordType RecordType, value string, ttl int) *Record` yang mengembalikan pointer ke struct `Record` baru.

Expected output ketika `go run main.go`:
```
Record created: example.com (A -> 93.184.216.34, TTL: 300s)
Alias created: www.example.com (CNAME -> example.com, TTL: 3600s)
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Dengan mendefinisikan custom type dan konstanta untuk tipe record DNS, kita menghindari kesalahan *typo* string dan mempermudah validasi. Pointer `*Record` akan menjadi fondasi utama dalam penyimpanan zone dan cache kita.
