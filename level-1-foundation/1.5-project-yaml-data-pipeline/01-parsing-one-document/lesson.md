# Parsing One Document

Selamat datang di Course 1.5: Project YAML Data Pipeline! Di course ini, kita akan membangun engine pemrosesan data in-memory menggunakan Go.
Format YAML sangat umum digunakan untuk file konfigurasi, data export, dan pertukaran log. Pada pelajaran pertama ini, kita akan mempelajari cara melakukan *unmarshaling* (decoding) satu dokumen YAML ke dalam struct Go menggunakan struct tag `yaml:"..."`.

## Contoh
example.go
```go
package main

import (
	"gopkg.in/yaml.v3"
)

type Product struct {
	ID       string  `yaml:"id"`
	Name     string  `yaml:"name"`
	Category string  `yaml:"category"`
	Price    float64 `yaml:"price"`
	Quantity int     `yaml:"quantity"`
}

func parseProduct(data []byte) (*Product, error) {
	var p Product
	err := yaml.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
```

## Penjelasan
- `gopkg.in/yaml.v3`: library pihak ketiga resmi dan standar industri di Go untuk mem-parse format YAML (sudah terpasang di modul proyekmu).
- Struct tag `yaml:"id"`: memberi tahu decoder YAML agar memetakan key YAML bernama `id` ke field struct `ID`.
- `yaml.Unmarshal(data []byte, out any) error`: mem-parsing byte YAML dan mengisi field struct yang ditunjuk oleh pointer `out`.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.5-project-yaml-data-pipeline/01-parsing-one-document/main.go).

1. Definisikan struct `Row` yang merepresentasikan satu baris data transaksi penjualan dengan tag YAML:
   - `ID` (`string`) dengan tag `yaml:"id"`
   - `Item` (`string`) dengan tag `yaml:"item"`
   - `Category` (`string`) dengan tag `yaml:"category"`
   - `Price` (`float64`) dengan tag `yaml:"price"`
   - `Quantity` (`int`) dengan tag `yaml:"quantity"`
   - `Total` (`float64`) dengan tag `yaml:"total"`
2. Buat fungsi `ParseRow(data []byte) (*Row, error)`:
   - Deklarasikan variabel struct `r Row`.
   - Jalankan `yaml.Unmarshal(data, &r)`.
   - Jika terjadi error, kembalikan `nil, err`.
   - Jika sukses, kembalikan pointer `&r, nil`.

Expected output ketika `go run main.go`:
```
Parsed: TX-101 | Laptop (electronics) | $1200.00 x 2
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Struct tags seperti `yaml:"field_name"` menjembatani penamaan format file eksternal dengan konvensi PascalCase struct field di Go.
