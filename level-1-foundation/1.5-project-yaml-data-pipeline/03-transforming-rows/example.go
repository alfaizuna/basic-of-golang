//go:build ignore

package main

import "strings"

// example.go — hanya untuk referensi, jangan diubah

type Row struct {
	ID       string  `yaml:"id"`
	Item     string  `yaml:"item"`
	Category string  `yaml:"category"`
	Price    float64 `yaml:"price"`
	Quantity int     `yaml:"quantity"`
	Total    float64 `yaml:"total"`
}

func CalculateTotals(rows []Row) []Row {
	result := make([]Row, len(rows))
	copy(result, rows)

	for i := range result {
		result[i].Total = result[i].Price * float64(result[i].Quantity)
	}
	return result
}

func NormalizeCategories(rows []Row) []Row {
	result := make([]Row, len(rows))
	copy(result, rows)

	for i := range result {
		result[i].Category = strings.ToLower(strings.TrimSpace(result[i].Category))
	}
	return result
}
