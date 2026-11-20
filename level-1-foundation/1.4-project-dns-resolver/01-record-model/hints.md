## Hint 1 (arah umum)
Definisikan tipe `type RecordType string`.
Gunakan blok `const (...)` untuk mendeklarasikan `TypeA` dan `TypeCNAME`.

---

## Hint 2 (sintaks yang relevan)
```go
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
```

---

## Hint 3 (hampir jawaban)
```go
func NewRecord(name string, recordType RecordType, value string, ttl int) *Record {
	return &Record{
		Name:  name,
		Type:  recordType,
		Value: value,
		TTL:   ttl,
	}
}
```
Buka komentar di dalam `main()` untuk mencetak output yang diharapkan.
