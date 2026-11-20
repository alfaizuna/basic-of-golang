## Hint 1 (arah umum)
Gunakan `decoder := yaml.NewDecoder(strings.NewReader(yamlData))`.
Deklarasikan slice penampung `var rows []Row`.

---

## Hint 2 (sintaks yang relevan)
Gunakan perulangan `for { ... }`:
```go
for {
    var r Row
    err := decoder.Decode(&r)
    if errors.Is(err, io.EOF) {
        break
    }
    if err != nil {
        return nil, err
    }
    rows = append(rows, r)
}
```

---

## Hint 3 (hampir jawaban)
```go
func ParseMultiDoc(yamlData string) ([]Row, error) {
	decoder := yaml.NewDecoder(strings.NewReader(yamlData))
	var rows []Row

	for {
		var r Row
		err := decoder.Decode(&r)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		rows = append(rows, r)
	}

	return rows, nil
}
```
