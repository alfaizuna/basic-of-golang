//go:build ignore

package main

// example.go — hanya untuk referensi, jangan diubah

type Row struct {
	ID       string  `yaml:"id"`
	Item     string  `yaml:"item"`
	Category string  `yaml:"category"`
	Price    float64 `yaml:"price"`
	Quantity int     `yaml:"quantity"`
	Total    float64 `yaml:"total"`
}

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
