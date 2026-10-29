### Hint 1
Jangan lupa menambahkan `"errors"` pada bagian `import`:
```go
import (
	"errors"
	"fmt"
	"strings"
)
```

### Hint 2
Bentuk fungsi `GenerateSlug`:
```go
func GenerateSlug(title string) (string, error) {
	if strings.TrimSpace(title) == "" {
		return "", errors.New("input text cannot be empty")
	}
	slug := transformToSlug(title)
	if slug == "" {
		return "", errors.New("no valid slug characters found")
	}
	return slug, nil
}
```

### Hint 3
Di dalam fungsi `main()`:
Panggil `GenerateSlug` untuk `title1` dan `title2`, lakukan pengecekan error dengan pola `if err != nil`, dan cetak output sesuai petunjuk pada `lesson.md`.
