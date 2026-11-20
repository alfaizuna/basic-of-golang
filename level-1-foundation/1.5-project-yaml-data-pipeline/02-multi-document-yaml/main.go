package main

import (
	"fmt"
	// "errors"
	// "io"
	// "strings"
	// "gopkg.in/yaml.v3"
)

type Row struct {
	ID       string  `yaml:"id"`
	Item     string  `yaml:"item"`
	Category string  `yaml:"category"`
	Price    float64 `yaml:"price"`
	Quantity int     `yaml:"quantity"`
	Total    float64 `yaml:"total"`
}

// TODO: Buat fungsi ParseMultiDoc(yamlData string) ([]Row, error)
// Gunakan strings.NewReader dan yaml.NewDecoder.
// Decode berulang kali sampai errors.Is(err, io.EOF).
func ParseMultiDoc(yamlData string) ([]Row, error) {
	return nil, nil
}

func main() {
	multiYAML := `
id: TX-101
item: Laptop
category: electronics
price: 1200.00
quantity: 2
---
id: TX-102
item: Desk
category: furniture
price: 450.00
quantity: 1
---
id: TX-103
item: Monitor
category: electronics
price: 300.00
quantity: 2
`
	rows, err := ParseMultiDoc(multiYAML)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Loaded %d rows from multi-document YAML\n", len(rows))
	for i, r := range rows {
		fmt.Printf("Row %d: %s (%s)\n", i+1, r.ID, r.Item)
	}
}
