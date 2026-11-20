## Hint 1 (arah umum)
Definisikan struct `Zone` dengan dua field: `Name string` dan `records map[string][]*Record`.
Di fungsi `NewZone`, alokasikan map dengan `make(map[string][]*Record)`.

---

## Hint 2 (sintaks yang relevan)
```go
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
```

---

## Hint 3 (hampir jawaban)
```go
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
