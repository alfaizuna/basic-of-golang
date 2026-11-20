# Hints

### Hint 1
Buat channel: `ch := make(chan int, len(values))`.

### Hint 2
Untuk setiap `v` di `values`, jalankan goroutine: `go func(val int) { ch <- val }(v)`.

### Hint 3
Lakukan loop `for i := 0; i < len(values); i++ { total += <-ch }` dan kembalikan total.
