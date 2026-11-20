## Hint 1 (arah umum)
Untuk `AddRecord`:
Periksa apakah `z == nil` atau `rec == nil`. Jika ya, langsung `return`.
Tambahkan pointer `rec` ke slice `z.records[rec.Name]`.

---

## Hint 2 (sintaks yang relevan)
```go
func (z *Zone) AddRecord(rec *Record) {
	if z == nil || rec == nil {
		return
	}
	z.records[rec.Name] = append(z.records[rec.Name], rec)
}
```

---

## Hint 3 (hampir jawaban)
```go
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
