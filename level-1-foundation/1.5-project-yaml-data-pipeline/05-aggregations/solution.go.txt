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

func main() {
	sampleRows := []Row{
		{ID: "TX-101", Item: "Laptop", Category: "electronics", Total: 2400.00},
		{ID: "TX-102", Item: "Desk", Category: "furniture", Total: 450.00},
		{ID: "TX-103", Item: "Monitor", Category: "electronics", Total: 600.00},
	}

	sum := Summarize(sampleRows)
	fmt.Printf("Summary: %d items, Total: $%.2f, Avg: $%.2f\n",
		sum.Count, sum.SumTotal, sum.AvgTotal)

	groups := GroupByCategory(sampleRows)
	fmt.Printf("Categories found: %d (electronics: %d, furniture: %d)\n",
		len(groups), len(groups["electronics"]), len(groups["furniture"]))
}
