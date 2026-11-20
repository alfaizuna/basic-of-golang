# Aggregations

Setelah data difilter dan dibersihkan, tujuan analitik utama dari data pipeline adalah menghitung ringkasan statistik (agregasi):
- Total pendapatan (`SumTotal`)
- Rata-rata transaksi (`AvgTotal`)
- Pengelompokan data transaksi berdasarkan kategori (*Group-by Category*)

## Contoh
example.go
```go
package main

type Summary struct {
	Count    int
	SumTotal float64
	AvgTotal float64
}

func Summarize(rows []Row) Summary {
	if len(rows) == 0 {
		return Summary{}
	}
	var sum float64
	for _, r := range rows {
		sum += r.Total
	}
	return Summary{
		Count:    len(rows),
		SumTotal: sum,
		AvgTotal: sum / float64(len(rows)),
	}
}

func GroupByCategory(rows []Row) map[string][]Row {
	grouped := make(map[string][]Row)
	for _, r := range rows {
		grouped[r.Category] = append(grouped[r.Category], r)
	}
	return grouped
}
```

## Penjelasan
- `map[string][]Row`: struktur data idiomatis di Go untuk mengelompokkan data berdasarkan kunci unik (*group-by*). Slice di dalam map akan otomatis bertambah saat elemen ditambahkan dengan `append`.
- Pengecekan `len(rows) == 0` penting untuk mencegah panic akibat pembagian dengan nol (*division by zero*) saat menghitung rata-rata.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.5-project-yaml-data-pipeline/05-aggregations/main.go).

1. Definisikan struct `Summary`:
   - `Count` (`int`)
   - `SumTotal` (`float64`)
   - `AvgTotal` (`float64`)
2. Buat fungsi `Summarize(rows []Row) Summary`:
   - Jika slice `rows` kosong, kembalikan `Summary{}`.
   - Hitung total akumulasi `r.Total` dari seluruh baris.
   - Hitung rata-rata `AvgTotal = SumTotal / float64(Count)`.
   - Kembalikan struct `Summary`.
3. Buat fungsi `GroupByCategory(rows []Row) map[string][]Row`:
   - Kelompokkan baris berdasarkan field `Category`.
   - Kembalikan map hasil pengelompokan.

Expected output ketika `go run main.go`:
```
Summary: 3 items, Total: $3450.00, Avg: $1150.00
Categories found: 2 (electronics: 2, furniture: 1)
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Map of slices (`map[string][]T`) adalah struktur data Go yang sangat efisien dan ekspresif untuk melakukan operasi pengelompokan (*grouping*) data agregasi.
