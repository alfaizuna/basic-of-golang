## Hint 1 (arah umum)
Untuk `Get`:
Gunakan comma-ok idiom dari map: `fl, ok := s.flags[name]`.
Jika `ok` bernilai `true`, kembalikan `(fl, true)`. Jika tidak, kembalikan `(nil, false)`.

---

## Hint 2 (sintaks yang relevan)
Untuk `FindOrCreate`:
Pertama gunakan `s.Get(name)` atau langsung cek map `s.flags[name]`.
Jika ditemukan, return pointernya.
Jika tidak, buat struct literal baru dengan `&Flag{...}`:
```go
fl := &Flag{
    Name:    name,
    Enabled: false,
    Rollout: 0,
}
s.flags[name] = fl
return fl
```

---

## Hint 3 (hampir jawaban)
```go
func (s *Store) Get(name string) (*Flag, bool) {
	fl, ok := s.flags[name]
	if !ok {
		return nil, false
	}
	return fl, true
}

func (s *Store) FindOrCreate(name string) *Flag {
	if fl, exists := s.Get(name); exists {
		return fl
	}
	fl := &Flag{
		Name:    name,
		Enabled: false,
		Rollout: 0,
	}
	s.flags[name] = fl
	return fl
}
```
