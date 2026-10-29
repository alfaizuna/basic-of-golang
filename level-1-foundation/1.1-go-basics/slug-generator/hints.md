### Hint 1
Awali fungsi `transformToSlug` dengan memanggil `cleanText(text)`. Kemudian lakukan perulangan pada string hasil pembersihan tersebut.

### Hint 2
Bentuk fungsi `transformToSlug`:
```go
func transformToSlug(text string) string {
	cleaned := cleanText(text)
	var result string
	inSeparator := false

	for _, r := range cleaned {
		if r == ' ' || r == '-' {
			if !inSeparator && len(result) > 0 {
				result += "-"
				inSeparator = true
			}
		} else {
			result += string(r)
			inSeparator = false
		}
	}
	return strings.Trim(result, "-")
}
```

### Hint 3
Di dalam fungsi `main()`:
```go
input := "  Go 101: Belajar #1!  "
slug := transformToSlug(input)
fmt.Printf("Input: %q\n", input)
fmt.Printf("Slug:  %q\n", slug)
```
