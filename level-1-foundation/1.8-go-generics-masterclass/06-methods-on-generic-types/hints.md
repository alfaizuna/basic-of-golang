# Hints

### Hint 1
Deklarasikan `type Stack[T any] struct { items []T }`.

### Hint 2
Di method Pop(): periksa `if len(s.items) == 0 { var zero T; return zero, false }`.

### Hint 3
Ambil `val := s.items[len(s.items)-1]`, potong `s.items = s.items[:len(s.items)-1]`, lalu return `val, true`.
