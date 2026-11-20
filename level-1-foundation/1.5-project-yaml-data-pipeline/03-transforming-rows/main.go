package main

import (
	"fmt"
	// "strings"
)

type Row struct {
	ID       string  `yaml:"id"`
	Item     string  `yaml:"item"`
	Category string  `yaml:"category"`
	Price    float64 `yaml:"price"`
	Quantity int     `yaml:"quantity"`
	Total    float64 `yaml:"total"`
}

// TODO 1: Buat fungsi CalculateTotals(rows []Row) []Row
// Hitung Total = Price * float64(Quantity) untuk setiap baris
func CalculateTotals(rows []Row) []Row {
	return nil
}

// TODO 2: Buat fungsi NormalizeCategories(rows []Row) []Row
// Ubah Category menjadi lowercase dan buang spasi berlebih
func NormalizeCategories(rows []Row) []Row {
	return nil
}

func main() {
	rawRows := []Row{
		{ID: "TX-101", Item: "Laptop", Category: "  ELECTRONICS ", Price: 1200.00, Quantity: 2},
		{ID: "TX-102", Item: "Desk", Category: "Furniture  ", Price: 450.00, Quantity: 1},
	}

	normalized := NormalizeCategories(rawRows)
	calculated := CalculateTotals(normalized)

	for _, r := range calculated {
		fmt.Printf("Transformed: %s | %s | %s | Total: $%.2f\n",
			r.ID, r.Item, r.Category, r.Total)
	}
}
