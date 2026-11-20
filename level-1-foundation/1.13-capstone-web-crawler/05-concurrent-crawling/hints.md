# Hints

### Hint 1
Deklarasikan `var wg sync.WaitGroup`.

### Hint 2
Untuk setiap `u := range urls`: `wg.Add(1)`, `go func(link string) { defer wg.Done(); processFn(link) }(u)`.

### Hint 3
Panggil `wg.Wait()`.
