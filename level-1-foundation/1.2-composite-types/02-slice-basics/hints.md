### Hint 1
Gunakan perulangan `for _, num := range numbers` untuk mengakses setiap angka di dalam slice tanpa perlu menggunakan variabel indeks.

### Hint 2
Bentuk fungsi `sumSlice`:
```go
func sumSlice(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
```

### Hint 3
Cukup tambahkan loop:
```go
for _, num := range numbers {
	total += num
}
```
sebelum `return total` di dalam fungsi `sumSlice`.
