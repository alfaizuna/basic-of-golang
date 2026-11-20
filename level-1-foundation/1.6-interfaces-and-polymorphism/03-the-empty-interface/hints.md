# Hints

### Hint 1
Gunakan loop `for _, item := range items` untuk menjelajahi slice `[]any`.

### Hint 2
Gunakan `fmt.Printf("[%T] %v\n", item, item)` untuk mencetak nama tipe (`%T`) dan nilainya (`%v`).

### Hint 3
Di `main()`, inisialisasi slice dengan:
```go
data := []any{100, "Go is awesome", true}
PrintAll(data)
```
