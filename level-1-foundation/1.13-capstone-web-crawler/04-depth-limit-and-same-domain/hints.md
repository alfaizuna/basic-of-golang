# Hints

### Hint 1
Periksa `if currentDepth > maxDepth { return false }`.

### Hint 2
Parse URL: `u, err := url.Parse(targetURL); if err != nil { return false }`.

### Hint 3
Kembalikan `u.Host == baseHost`.
