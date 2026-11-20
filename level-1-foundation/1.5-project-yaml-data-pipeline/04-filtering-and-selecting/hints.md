## Hint 1 (arah umum)
Definisikan `type Predicate func(Row) bool`.
Gunakan fungsi `Filter` sebagai pondasi untuk fungsi filter turunan lainnya.

---

## Hint 2 (sintaks yang relevan)
```go
func Filter(rows []Row, predicate Predicate) []Row {
	var filtered []Row
	for _, r := range rows {
		if predicate(r) {
			filtered = append(filtered, r)
		}
	}
	return filtered
}
```

---

## Hint 3 (hampir jawaban)
```go
func FilterByCategory(rows []Row, category string) []Row {
	return Filter(rows, func(r Row) bool {
		return r.Category == category
	})
}

func FilterMinTotal(rows []Row, minTotal float64) []Row {
	return Filter(rows, func(r Row) bool {
		return r.Total >= minTotal
	})
}
```
