### Hint 1
Go hanya memiliki kata kunci `for` untuk seluruh jenis perulangan. Gunakan operator modulo `%` untuk mengecek apakah bilangan ganjil (`i % 2 != 0`).

### Hint 2
Bentuk fungsi `sumEven`:
```go
func sumEven(limit int) int {
	total := 0
	for i := 1; i <= limit; i++ {
		if i%2 != 0 {
			continue
		}
		total += i
	}
	return total
}
```

### Hint 3
Bentuk fungsi `countUntil`:
```go
func countUntil(target int) int {
	count := 0
	steps := 0
	for {
		count += 3
		steps++
		if count >= target {
			break
		}
	}
	return steps
}
```
Lalu panggil `fmt.Println(sumEven(10))` dan `fmt.Println(countUntil(10))` di dalam `main()`.
