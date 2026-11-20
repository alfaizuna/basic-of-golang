## Hint 1 (arah umum)
Definisikan struct `Summary` dengan 3 field: `Count int`, `SumTotal float64`, dan `AvgTotal float64`.
Di fungsi `Summarize`, periksa apakah `len(rows) == 0`. Jika ya, langsung kembalikan `Summary{}`.

---

## Hint 2 (sintaks yang relevan)
Untuk `Summarize`:
```go
var sum float64
for _, r := range rows {
    sum += r.Total
}
return Summary{
    Count:    len(rows),
    SumTotal: sum,
    AvgTotal: sum / float64(len(rows)),
}
```

---

## Hint 3 (hampir jawaban)
```go
func GroupByCategory(rows []Row) map[string][]Row {
	grouped := make(map[string][]Row)
	for _, r := range rows {
		grouped[r.Category] = append(grouped[r.Category], r)
	}
	return grouped
}
```
