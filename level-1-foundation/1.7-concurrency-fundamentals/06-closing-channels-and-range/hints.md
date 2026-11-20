# Hints

### Hint 1
Buat channel: `ch := make(chan int, n)`.

### Hint 2
Di goroutine: gunakan loop `for i := 0; i < n; i++` dengan variabel Fibonacci `a, b = b, a+b`.

### Hint 3
Pastikan memanggil `close(ch)` setelah loop selesai di dalam goroutine.
