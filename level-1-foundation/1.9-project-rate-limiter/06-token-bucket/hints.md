# Hints

### Hint 1
Hitung `elapsed := tb.now().Sub(tb.lastRefill).Seconds()`.

### Hint 2
Tambahkan token: `tb.tokens += elapsed * tb.refillRate` lalu batasi ke `tb.capacity`.

### Hint 3
Kurangi token `tb.tokens -= 1.0` jika `tb.tokens >= 1.0`.
