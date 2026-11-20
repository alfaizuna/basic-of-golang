package main

import "fmt"

type Row struct {
	ID       string  `yaml:"id"`
	Item     string  `yaml:"item"`
	Category string  `yaml:"category"`
	Price    float64 `yaml:"price"`
	Quantity int     `yaml:"quantity"`
	Total    float64 `yaml:"total"`
}

// TODO 1: Definisikan tipe Predicate func(Row) bool
// type Predicate func(Row) bool

// TODO 2: Buat fungsi Filter(rows []Row, predicate func(Row) bool) []Row
func Filter(rows []Row, predicate func(Row) bool) []Row {
	return nil
}

// TODO 3: Buat fungsi FilterByCategory(rows []Row, category string) []Row
func FilterByCategory(rows []Row, category string) []Row {
	return nil
}

// TODO 4: Buat fungsi FilterMinTotal(rows []Row, minTotal float64) []Row
func FilterMinTotal(rows []Row, minTotal float64) []Row {
	return nil
}

func main() {
	sampleRows := []Row{
		{ID: "TX-101", Item: "Laptop", Category: "electronics", Total: 2400.00},
		{ID: "TX-102", Item: "Desk", Category: "furniture", Total: 450.00},
		{ID: "TX-103", Item: "Monitor", Category: "electronics", Total: 600.00},
	}

	electronics := FilterByCategory(sampleRows, "electronics")
	fmt.Printf("Filtered category 'electronics': %d rows\n", len(electronics))

	highValue := FilterMinTotal(sampleRows, 500.00)
	fmt.Printf("Filtered min total $500: %d rows\n", len(highValue))
}
