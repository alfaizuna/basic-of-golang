package main

import (
	"fmt"
	// "gopkg.in/yaml.v3"
)

// TODO 1: Definisikan struct Row dengan tag yaml:
// - ID (string)        `yaml:"id"`
// - Item (string)      `yaml:"item"`
// - Category (string)  `yaml:"category"`
// - Price (float64)    `yaml:"price"`
// - Quantity (int)     `yaml:"quantity"`
// - Total (float64)    `yaml:"total"`

type Row struct {
	// lengkapi di sini
}

// TODO 2: Buat fungsi ParseRow(data []byte) (*Row, error)
func ParseRow(data []byte) (*Row, error) {
	return nil, nil
}

func main() {
	sampleYAML := `
id: TX-101
item: Laptop
category: electronics
price: 1200.00
quantity: 2
`
	row, err := ParseRow([]byte(sampleYAML))
	if err != nil {
		fmt.Println("Error parsing:", err)
		return
	}
	if row != nil {
		fmt.Printf("Parsed: %s | %s (%s) | $%.2f x %d\n",
			row.ID, row.Item, row.Category, row.Price, row.Quantity)
	} else {
		fmt.Println("Row is nil")
	}
}
