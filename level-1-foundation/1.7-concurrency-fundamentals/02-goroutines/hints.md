# Hints

### Hint 1
Fungsi LaunchWorker dapat menggunakan `go func() { out <- fmt.Sprintf("Worker %d done", id) }()`.

### Hint 2
Di main, buat channel `out := make(chan string)`, panggil `LaunchWorker(1, out)`, lalu baca dengan `msg := <-out`.

### Hint 3
Pastikan Anda mencetak string yang diterima dari channel di main().
