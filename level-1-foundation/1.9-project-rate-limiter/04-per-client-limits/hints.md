# Hints

### Hint 1
Periksa `cs, ok := l.clients[clientKey]`. Jika `!ok`, buat baru `cs = &ClientState{windowStart: cur}` dan simpan di `l.clients[clientKey]`.

### Hint 2
Periksa `if cur.Sub(cs.windowStart) >= l.window { cs.windowStart = cur; cs.count = 0 }`.

### Hint 3
Jika `cs.count < l.limit { cs.count++; return true } else { return false }`.
