## Hint 1 (arah umum)
Buat set pelacak `visited := make(map[string]bool)` sebelum perulangan dimulai.
Gunakan variabel penghitung kedalaman `depth := 0`.

---

## Hint 2 (sintaks yang relevan)
```go
for {
    if depth >= maxCNAMEDepth {
        return nil, errors.New("maximum CNAME depth exceeded")
    }
    if visited[current] {
        return nil, errors.New("cname loop detected")
    }
    visited[current] = true
    // lanjutkan pencarian record A dan CNAME...
}
```

---

## Hint 3 (hampir jawaban)
```go
func (z *Zone) Resolve(name string) (*Record, error) {
	visited := make(map[string]bool)
	current := name
	depth := 0

	for {
		if depth >= maxCNAMEDepth {
			return nil, errors.New("maximum CNAME depth exceeded")
		}
		if visited[current] {
			return nil, errors.New("cname loop detected")
		}
		visited[current] = true

		if rec := z.Lookup(current, TypeA); rec != nil {
			return rec, nil
		}

		cname := z.Lookup(current, TypeCNAME)
		if cname == nil {
			return nil, errors.New("record not found")
		}

		current = cname.Value
		depth++
	}
}
```
