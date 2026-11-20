## Hint 1 (arah umum)
Definisikan struct `Flag` dengan 3 field: `Name string`, `Enabled bool`, dan `Rollout int`.
Definisikan struct `Store` dengan field `flags map[string]*Flag`.

---

## Hint 2 (sintaks yang relevan)
Untuk `NewStore`, inisialisasi map dengan `make(map[string]*Flag)`:
```go
func NewStore() *Store {
	return &Store{
		flags: make(map[string]*Flag),
	}
}
```

Untuk `Count`:
```go
func (s *Store) Count() int {
	if s == nil || s.flags == nil {
		return 0
	}
	return len(s.flags)
}
```

---

## Hint 3 (hampir jawaban)
```go
type Flag struct {
	Name    string
	Enabled bool
	Rollout int
}

type Store struct {
	flags map[string]*Flag
}

func NewStore() *Store {
	return &Store{
		flags: make(map[string]*Flag),
	}
}

func (s *Store) Count() int {
	if s == nil || s.flags == nil {
		return 0
	}
	return len(s.flags)
}
```
