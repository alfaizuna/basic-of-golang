## Hint 1 (arah umum)
Periksa apakah receiver `z == nil`. Jika ya, kembalikan `nil`.
Buat zone baru dengan `target := NewZone(z.Name)`.

---

## Hint 2 (sintaks yang relevan)
Iterasi seluruh record di `z.records`:
```go
for _, recs := range z.records {
    for _, r := range recs {
        // Buat struct baru, JANGAN masukkan r langsung!
        copiedRec := &Record{
            Name:  r.Name,
            Type:  r.Type,
            Value: r.Value,
            TTL:   r.TTL,
        }
        target.AddRecord(copiedRec)
    }
}
```

---

## Hint 3 (hampir jawaban)
```go
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
