# Hints

### Hint 1
Definisikan interface `RateLimiter`.

### Hint 2
Di method Allow pada `*SimulatedLimiter`: jika `s.current < s.maxRequests { s.current++; return true } return false`.

### Hint 3
Di `ExecuteRequests`: loop `for i := 0; i < count; i++ { if rl.Allow(key) { allowed++ } }`.
