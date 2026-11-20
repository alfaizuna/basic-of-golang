## Hint 1 (arah umum)
Buka komentar import `"gopkg.in/yaml.v3"` di bagian atas file.
Definisikan struct `Row` dengan 6 field sesuai spesifikasi, lengkap dengan tag `yaml:"..."` di setiap fieldnya.

---

## Hint 2 (sintaks yang relevan)
```go
type Row struct {
	ID       string  `yaml:"id"`
	Item     string  `yaml:"item"`
	Category string  `yaml:"category"`
	Price    float64 `yaml:"price"`
	Quantity int     `yaml:"quantity"`
	Total    float64 `yaml:"total"`
}
```

---

## Hint 3 (hampir jawaban)
```go
func ParseRow(data []byte) (*Row, error) {
	var r Row
	if err := yaml.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
```
