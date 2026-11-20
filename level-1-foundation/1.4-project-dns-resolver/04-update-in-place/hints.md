## Hint 1 (arah umum)
Untuk `UpdateValue` dan `UpdateTTL`:
Periksa apakah receiver `r == nil`. Jika tidak, ubah langsung `r.Value = newValue` atau `r.TTL = newTTL`.

---

## Hint 2 (sintaks yang relevan)
```go
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

---

## Hint 3 (hampir jawaban)
```go
func (z *Zone) UpdateRecord(name string, recordType RecordType, newValue string, newTTL int) bool {
	rec := z.Lookup(name, recordType)
	if rec == nil {
		return false
	}
	rec.UpdateValue(newValue)
	rec.UpdateTTL(newTTL)
	return true
}
```
