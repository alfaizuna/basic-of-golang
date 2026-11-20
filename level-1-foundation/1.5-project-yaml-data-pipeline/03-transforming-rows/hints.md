## Hint 1 (arah umum)
Buat salinan slice:
```go
res := make([]Row, len(rows))
copy(res, rows)
```
Kemudian iterasi lewat indeks `for i := range res { ... }`.

---

## Hint 2 (sintaks yang relevan)
Untuk `CalculateTotals`:
```go
res[i].Total = res[i].Price * float64(res[i].Quantity)
```

Untuk `NormalizeCategories`:
```go
res[i].Category = strings.ToLower(strings.TrimSpace(res[i].Category))
```

---

## Hint 3 (hampir jawaban)
```go
func CalculateTotals(rows []Row) []Row {
	res := make([]Row, len(rows))
	copy(res, rows)
	for i := range res {
		res[i].Total = res[i].Price * float64(res[i].Quantity)
	}
	return res
}

func NormalizeCategories(rows []Row) []Row {
	res := make([]Row, len(rows))
	copy(res, rows)
	for i := range res {
		res[i].Category = strings.ToLower(strings.TrimSpace(res[i].Category))
	}
	return res
}
```
