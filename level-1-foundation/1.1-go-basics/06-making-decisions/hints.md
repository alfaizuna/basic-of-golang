### Hint 1
Awali dengan syntax if initializer: `if score := 85; score >= 90 { ... }`.

### Hint 2
Lanjutkan dengan `} else if score >= 80 { ... }`, lalu `} else if score >= 70 { ... }`, dan diakhiri `} else { ... }`.
Ingat: `else` harus berada di baris yang sama persis dengan `}` penutup blok sebelumnya.

### Hint 3
Lengkapi `func main()` dengan kode berikut:
```go
if score := 85; score >= 90 {
	fmt.Println("Grade: A")
} else if score >= 80 {
	fmt.Println("Grade: B")
} else if score >= 70 {
	fmt.Println("Grade: C")
} else {
	fmt.Println("Grade: F")
}
```
