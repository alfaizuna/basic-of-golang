### Hint 1
Deklarasikan `const MaxScore = 100` di luar `main()`. Buat `StatusDraft` dan `StatusPublished` dalam satu blok `const (...)` menggunakan `iota`.

### Hint 2
Untuk `iota`, tulis:
```go
const (
	StatusDraft = iota
	StatusPublished
)
```
Dan untuk zero value di dalam `main()`, cukup gunakan `var defaultAttempts int` dan `var isReady bool` tanpa tanda sama dengan `=`.

### Hint 3
Lengkapi `main.go` dengan:
```go
const MaxScore = 100

const (
	StatusDraft = iota
	StatusPublished
)

func main() {
	var defaultAttempts int
	var isReady bool

	fmt.Println(MaxScore)
	fmt.Println(StatusDraft)
	fmt.Println(StatusPublished)
	fmt.Println(defaultAttempts)
	fmt.Println(isReady)
}
```
