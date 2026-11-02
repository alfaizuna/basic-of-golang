### Hint 1
Di dalam `countWords`, inisialisasi map dengan `result := make(map[string]int)`.

### Hint 2
Lakukan iterasi kata menggunakan `for _, w := range words` lalu gunakan `result[w]++`. Terakhir, `return result`.

### Hint 3
Kerangka kode lengkap:
```go
func countWords(words []string) map[string]int {
	result := make(map[string]int)
	for _, w := range words {
		result[w]++
	}
	return result
}

func main() {
	words := []string{"go", "rust", "go", "python", "go", "rust"}
	counts := countWords(words)

	fmt.Println("Go:", counts["go"])
	fmt.Println("Rust:", counts["rust"])
	fmt.Println("Python:", counts["python"])
}
```
