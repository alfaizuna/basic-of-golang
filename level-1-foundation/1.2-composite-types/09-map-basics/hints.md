### Hint 1
Gunakan map literal `map[string]int{...}` untuk membuat map awal, dan gunakan sintaks assignment `scores["Charlie"] = 95` untuk menambah elemen.

### Hint 2
Untuk menghapus, panggil fungsi `delete`:
```go
delete(scores, "Bob")
```

### Hint 3
Kerangka kode lengkap:
```go
scores := map[string]int{"Alice": 90, "Bob": 85}
scores["Charlie"] = 95
delete(scores, "Bob")

fmt.Println("Alice:", scores["Alice"])
fmt.Println("Total:", len(scores))
```
