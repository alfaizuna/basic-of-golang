# Hints

### Hint 1
Kunci dengan `s.mu.Lock()` dan `defer s.mu.Unlock()`.

### Hint 2
Periksa `if s.visited[url] { return false }` lalu setel `s.visited[url] = true` dan return `true`.
