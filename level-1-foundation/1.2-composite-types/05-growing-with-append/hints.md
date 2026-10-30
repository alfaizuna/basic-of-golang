### Hint 1
Gunakan fungsi bawaan `append(scores, newScore)` di dalam fungsi `addScore` dan kembalikan nilainya dengan kata kunci `return`. Di dalam `main()`, tampung kembali hasil pemanggilan fungsinya ke variabel `scores`.

### Hint 2
Di dalam `addScore`:
```go
return append(scores, newScore)
```
Di dalam `main()`:
```go
scores = addScore(scores, 100)
```

### Hint 3
Kerangka lengkap kode:
```go
func addScore(scores []int, newScore int) []int {
	return append(scores, newScore)
}

func main() {
	scores := []int{80, 90}
	scores = addScore(scores, 100)
	scores = addScore(scores, 95)

	fmt.Println("Scores:", scores)
	fmt.Println("Length:", len(scores))
}
```
