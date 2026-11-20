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

type Predicate func(Row) bool

func Filter(rows []Row, keep Predicate) []Row {
	var filtered []Row
	for _, r := range rows {
		if keep(r) {
			filtered = append(filtered, r)
		}
	}
	return filtered
}
