# Hints

### Hint 1
Sertakan `mu sync.Mutex` di SafeRateLimiter.

### Hint 2
Di awal Allow() dan GetCount(), panggil `s.mu.Lock()` lalu `defer s.mu.Unlock()`.
