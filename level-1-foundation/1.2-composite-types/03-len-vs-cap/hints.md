### Hint 1
Fungsi `getLenAndCap` harus mengembalikan dua nilai integer sekaligus:
`return len(s), cap(s)`

### Hint 2
Untuk membuat sub-slice dari indeks 1 sampai sebelum 3:
`part := data[1:3]`

### Hint 3
Di dalam `main()`:
```go
part := data[1:3]
l, c := getLenAndCap(part)
fmt.Println(l, c)
```
