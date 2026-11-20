## Hint 1 (arah umum)
Gunakan variabel `current := name` dan loop `for { ... }`.
Di setiap iterasi, cek dulu apakah ada record `TypeA`.

---

## Hint 2 (sintaks yang relevan)
```go
current := name
for {
    if rec := z.Lookup(current, TypeA); rec != nil {
        return rec, nil
    }
    // Jika tidak ada A, cari CNAME...
}
```

---

## Hint 3 (hampir jawaban)
```go
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
