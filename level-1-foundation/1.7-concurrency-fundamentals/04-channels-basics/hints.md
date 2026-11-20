# Hints

### Hint 1
Buat channel dengan `ch := make(chan int)`.

### Hint 2
Luncurkan goroutine dengan `go func() { ch <- n * n }()`.

### Hint 3
Di luar goroutine, terima nilainya dengan `return <-ch`.
