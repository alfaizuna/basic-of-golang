### Hint 1
Gunakan loop `for k := range scores` dan tambahkan setiap `k` ke slice `keys` menggunakan `append`.

### Hint 2
Panggil fungsi sorting:
```go
sort.Strings(keys)
```
Kemudian lakukan iterasi pada `keys`:
```go
for _, k := range keys {
    fmt.Printf("%s: %d\n", k, scores[k])
}
```

### Hint 3
Kerangka kode lengkap:
```go
scores := map[string]int{"Charlie": 75, "Alice": 95, "Bob": 85}

var keys []string
for k := range scores {
	keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
	fmt.Printf("%s: %d\n", k, scores[k])
}
```
