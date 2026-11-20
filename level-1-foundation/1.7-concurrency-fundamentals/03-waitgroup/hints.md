# Hints

### Hint 1
Deklarasikan `var wg sync.WaitGroup` dan loop dari `0` sampai `n`.

### Hint 2
Di setiap iterasi loop, panggil `wg.Add(1)` lalu jalankan `go func() { defer wg.Done(); mu.Lock(); total += 10; mu.Unlock() }()`.

### Hint 3
Setelah loop selesai, panggil `wg.Wait()` dan kembalikan `total`.
