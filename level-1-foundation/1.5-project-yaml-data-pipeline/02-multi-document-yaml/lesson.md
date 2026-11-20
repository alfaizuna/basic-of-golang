# Multi-Document YAML

Format YAML mendukung penyimpanan beberapa dokumen di dalam satu file atau teks yang sama dengan memisahkannya menggunakan tanda tiga strip (`---`).
Contohnya:
```yaml
id: TX-101
item: Laptop
---
id: TX-102
item: Desk
```

Untuk membaca data multi-dokumen secara efisien, kita menggunakan `yaml.NewDecoder(reader)` yang dipanggil berulang kali di dalam loop sampai menemui error akhir file: `io.EOF`.

## Contoh
example.go
```go
package main

import (
	"errors"
	"io"
	"strings"
	"gopkg.in/yaml.v3"
)

func parseAllRows(yamlContent string) ([]Row, error) {
	decoder := yaml.NewDecoder(strings.NewReader(yamlContent))
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

## Penjelasan
- `yaml.NewDecoder(r io.Reader)`: membuat instance decoder streaming untuk membaca data YAML berurutan.
- `decoder.Decode(&r)`: mem-parsing satu dokumen YAML berikutnya ke variabel `r`.
- `errors.Is(err, io.EOF)`: tanda penanda dari Go bahwa seluruh aliran dokumen telah selesai dibaca tanpa error.
- Pola streaming decoder ini jauh lebih hemat memori dan aman daripada memecah string dokumen secara manual dengan split string.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.5-project-yaml-data-pipeline/02-multi-document-yaml/main.go).

Buat fungsi `ParseMultiDoc(yamlData string) ([]Row, error)`:
1. Gunakan `strings.NewReader(yamlData)` dan `yaml.NewDecoder(...)`.
2. Lakukan iterasi perulangan:
   - Deklarasikan variabel `var r Row`.
   - Panggil `err := dec.Decode(&r)`.
   - Jika `errors.Is(err, io.EOF)`, hentikan perulangan (`break`).
   - Jika `err != nil`, kembalikan `nil, err`.
   - Tambahkan `r` ke dalam slice `rows`.
3. Kembalikan `rows, nil`.

Expected output ketika `go run main.go`:
```
Loaded 3 rows from multi-document YAML
Row 1: TX-101 (Laptop)
Row 2: TX-102 (Desk)
Row 3: TX-103 (Monitor)
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Gunakan `yaml.NewDecoder` dengan loop dan `errors.Is(err, io.EOF)` untuk memproses multi-dokumen YAML secara streaming dan idiomatik.
