# Hints

### Hint 1
Deklarasikan `func SafeRun(fn func()) (err error)`.

### Hint 2
Di awal fungsi: `defer func() { if r := recover(); r != nil { err = fmt.Errorf("panic recovered: %v", r) } }()`.

### Hint 3
Panggil `fn()`.
