package main

import (
	"fmt"
	"strings"
)

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

func NormalizeCategories(rows []Row) []Row {
	res := make([]Row, len(rows))
	copy(res, rows)
	for i := range res {
		res[i].Category = strings.ToLower(strings.TrimSpace(res[i].Category))
	}
	return res
}

func CalculateTotals(rows []Row) []Row {
	res := make([]Row, len(rows))
	copy(res, rows)
	for i := range res {
		res[i].Total = res[i].Price * float64(res[i].Quantity)
	}
	return res
}

func FilterMinTotal(min float64) Stage {
	return func(rows []Row) []Row {
		var out []Row
		for _, r := range rows {
			if r.Total >= min {
				out = append(out, r)
			}
		}
		return out
	}
}

// TODO 1: Definisikan tipe Stage func([]Row) []Row
// type Stage func([]Row) []Row

// TODO 2: Definisikan struct Pipeline dengan field stages ([]Stage)
type Pipeline struct {
	// lengkapi di sini
}

// TODO 3: Buat fungsi NewPipeline(stages ...Stage) *Pipeline
func NewPipeline(stages ...Stage) *Pipeline {
	return nil
}

// TODO 4: Buat method (p *Pipeline) AddStage(stage Stage)
func (p *Pipeline) AddStage(stage Stage) {
}

// TODO 5: Buat method (p *Pipeline) Execute(input []Row) []Row
func (p *Pipeline) Execute(input []Row) []Row {
	return nil
}

func main() {
	sampleRows := []Row{
		{ID: "TX-101", Item: "Laptop", Category: "  ELECTRONICS ", Price: 1200.00, Quantity: 2},
		{ID: "TX-102", Item: "Desk", Category: "furniture", Price: 450.00, Quantity: 1},
		{ID: "TX-103", Item: "Mouse", Category: "ELECTRONICS", Price: 25.00, Quantity: 2},
	}

	pipeline := NewPipeline(
		NormalizeCategories,
		CalculateTotals,
	)
	if pipeline != nil {
		// Tambahkan stage filter belanjaan >= $1000
		pipeline.AddStage(FilterMinTotal(1000.00))

		processed := pipeline.Execute(sampleRows)
		fmt.Printf("Pipeline executed: %d raw -> %d processed row\n", len(sampleRows), len(processed))
		if len(processed) > 0 {
			r := processed[0]
			fmt.Printf("Final row: %s | %s | %s | $%.2f\n", r.ID, r.Item, r.Category, r.Total)
		}

		sum := Summarize(processed)
		fmt.Printf("Pipeline Summary: %d items, Total: $%.2f, Avg: $%.2f\n",
			sum.Count, sum.SumTotal, sum.AvgTotal)
	}
}
