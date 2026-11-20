# Hints

### Hint 1
Di method Inc(): panggil `c.mu.Lock()`, `defer c.mu.Unlock()`, dan `c.value++`.

### Hint 2
Di method Value(): panggil `c.mu.Lock()`, `defer c.mu.Unlock()`, dan `return c.value`.

### Hint 3
Pastikan kedua method menggunakan pointer receiver `*SafeCounter`.
