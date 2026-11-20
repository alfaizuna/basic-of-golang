# Hints

### Hint 1
Buat channel dengan `ch := make(chan string, len(jobs))`.

### Hint 2
Gunakan loop `for _, job := range jobs { ch <- job }`.

### Hint 3
Kembalikan `ch, cap(ch)`.
