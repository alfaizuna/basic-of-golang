# Hints

### Hint 1
Periksa kondisi reset dengan `if cur.Sub(l.windowStart) >= l.window`.

### Hint 2
Di dalam blok reset, update `l.windowStart = cur` dan `l.count = 0`.

### Hint 3
Lalu evaluasi `if l.count < l.limit` seperti biasa.
