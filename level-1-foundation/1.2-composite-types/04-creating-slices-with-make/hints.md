### Hint 1
Gunakan fungsi bawaan `make` dengan tiga argumen: tipe slice, panjang, dan kapasitas:
`return make([]int, length, capacity)`

### Hint 2
Untuk mengisi nilai pada indeks elemen slice yang sudah memiliki panjang (`len >= 2`):
```go
s[0] = 100
s[1] = 200
```

### Hint 3
Di dalam `main()`:
```go
s := createBufferedSlice(2, 5)
s[0] = 100
s[1] = 200
fmt.Println(s)
fmt.Println(len(s), cap(s))
```
