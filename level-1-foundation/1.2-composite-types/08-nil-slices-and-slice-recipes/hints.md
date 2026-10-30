### Hint 1
Untuk membuang elemen di posisi `index`, gabungkan potongan slice sebelum elemen tersebut (`items[:index]`) dengan potongan slice setelahnya (`items[index+1:]`) menggunakan fungsi `append`.

### Hint 2
Ingat bahwa slice kedua harus di-unpack menggunakan tanda tiga titik `...`:
```go
return append(items[:index], items[index+1:]...)
```

### Hint 3
Kerangka kode lengkap:
```go
func removeAt(items []string, index int) []string {
	return append(items[:index], items[index+1:]...)
}

func main() {
	languages := []string{"Go", "Java", "Python", "Ruby"}
	languages = removeAt(languages, 1)

	fmt.Println("Languages:", languages)

	var empty []string
	fmt.Println("Is Nil:", empty == nil)
}
```
