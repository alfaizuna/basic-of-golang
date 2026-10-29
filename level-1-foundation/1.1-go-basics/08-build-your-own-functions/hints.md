### Hint 1
Deklarasikan fungsi `greet` dan `double` di luar fungsi `main()`. Ingat: sebutkan tipe parameter setelah nama parameternya, dan tipe return sebelum kurung kurawal `{`.

### Hint 2
Bentuk fungsi `greet` dan `double`:
```go
func greet(name string) string {
	return "Halo, " + name
}

func double(n int) int {
	return n * 2
}
```

### Hint 3
Di dalam `func main()`, panggil kedua fungsi tersebut dan teruskan ke `fmt.Println`:
```go
fmt.Println(greet("Go"))
fmt.Println(double(21))
```
