# Hints

### Hint 1
Periksa `if now == nil { now = time.Now }` di NewLimiter.

### Hint 2
Kembalikan pointer `&Limiter{Limit: limit, Now: now}`.

### Hint 3
Method Allow: `return l.Limit > 0`.
