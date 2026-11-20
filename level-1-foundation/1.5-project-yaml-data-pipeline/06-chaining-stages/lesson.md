# Chaining Stages

Tahap penutup Course 1.5: merangkai (*chaining*) semua fungsi pemrosesan data menjadi sebuah satu kesatuan **Pipeline** yang modular dan dapat dieksekusi secara berurutan.

Setiap tahap memiliki bentuk fungsi yang seragam: menerima sekumpulan baris dan mengembalikan sekumpulan baris baru:
```go
type Stage func([]Row) []Row
```

Dengan bentuk ini, kita dapat merangkai tahap-tahap tersebut ke dalam struct `Pipeline` yang mengeksekusinya secara berurutan:
`Raw Rows` ➡️ `Normalize` ➡️ `Calculate Totals` ➡️ `Filter` ➡️ `Summary Output`.

## Contoh
example.go
```go
package main

type Stage func([]Row) []Row

type Pipeline struct {
	stages []Stage
}

func NewPipeline(stages ...Stage) *Pipeline {
	return &Pipeline{stages: stages}
}

func (p *Pipeline) Execute(initial []Row) []Row {
	current := initial
	for _, stage := range p.stages {
		current = stage(current)
	}
	return current
}
```

## Penjelasan
- `Stage func([]Row) []Row`: tipe fungsi perantara (*pipeline stage*).
- `stages ...Stage`: parameter variadik yang memungkinkan pendaftaran tahapan sebanyak apapun secara deklaratif.
- `current = stage(current)`: output dari satu tahap otomatis menjadi input untuk tahap berikutnya.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.5-project-yaml-data-pipeline/06-chaining-stages/main.go).

1. Definisikan tipe fungsi `Stage`:
   ```go
   type Stage func([]Row) []Row
   ```
2. Definisikan struct `Pipeline`:
   - `stages` (`[]Stage`)
3. Buat konstruktor `NewPipeline(stages ...Stage) *Pipeline`:
   - Mengembalikan pointer ke `Pipeline` baru.
4. Buat method `(p *Pipeline) AddStage(stage Stage)`:
   - Tambahkan `stage` ke slice `p.stages`.
5. Buat method `(p *Pipeline) Execute(input []Row) []Row`:
   - Alirkan data input melalui setiap tahap di `p.stages`.
   - Kembalikan hasil akhir setelah melewati seluruh tahap.

Expected output ketika `go run main.go`:
```
Pipeline executed: 3 raw -> 1 processed row
Final row: TX-101 | Laptop | electronics | $2400.00
Pipeline Summary: 1 items, Total: $2400.00, Avg: $2400.00
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Selamat! Kamu telah menyelesaikan Course 1.5: Project YAML Data Pipeline!
Dengan merangkai fungsi (*function chaining*) bertipe seragam `Stage func([]Row) []Row`, kamu berhasil membangun arsitektur pipeline data yang sangat modular, mudah dites, dan dapat diperluas tanpa batas.
