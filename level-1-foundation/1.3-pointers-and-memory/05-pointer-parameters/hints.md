### Hint 1
Di dalam fungsi `double(n *int)`, periksa `if n != nil` terlebih dahulu. Ini adalah praktik keamanan pointer (guard) agar program tidak crash jika pointer kosong.

### Hint 2
- Untuk mendereference dan mengalikan nilai: `*n = *n * 2` atau `*n *= 2`.
- Di dalam `main()`, oper alamat dari `score` menggunakan tanda ampersand: `double(&score)`.

### Hint 3
Kerangka kode lengkapnya:
```go
func double(n *int) {
	if n != nil {
		*n *= 2
	}
}

func main() {
	score := 25
	double(&score)
	fmt.Println("Score setelah didouble:", score)
}
```
