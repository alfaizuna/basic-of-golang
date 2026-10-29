### Hint 1
Untuk menghitung byte, gunakan fungsi bawaan `len(text)`. Untuk menghitung karakter (rune) Unicode sesungguhnya, gunakan fungsi `utf8.RuneCountInString(text)`.

### Hint 2
Untuk `normalizeText`, kamu bisa merangkai dua fungsi paket `strings`:
```go
func normalizeText(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}
```

### Hint 3
Di dalam fungsi `main()`:
```go
b, r := countBytesAndRunes("Go 🚀")
fmt.Println(b, r)

cleaned := normalizeText("  Belajar GOLANG  ")
fmt.Println(cleaned)
```
