### Hint 1
Untuk mengembalikan hasil sukses dari fungsi `divide(a, b float64) (float64, error)`, ganti `return 0, nil` menjadi `return a / b, nil`. Nilai `nil` menandakan bahwa operasi berhasil tanpa error.

### Hint 2
Pola penanganan error untuk pemanggilan kedua (`divide(10, 0)`):
```go
res2, err2 := divide(10, 0)
if err2 != nil {
    fmt.Println(err2)
} else {
    fmt.Println(res2)
}
```

### Hint 3
Di dalam fungsi `divide`:
```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
```
Lalu di dalam `main()`, cetak `res1` pada blok `else` panggilan pertama, dan ulangi pola yang sama untuk `divide(10, 0)`.
