//go:build ignore

package main

import (
	"gopkg.in/yaml.v3"
)

// example.go — hanya untuk referensi, jangan diubah

type Row struct {
	ID       string  `yaml:"id"`
	Item     string  `yaml:"item"`
	Category string  `yaml:"category"`
	Price    float64 `yaml:"price"`
	Quantity int     `yaml:"quantity"`
	Total    float64 `yaml:"total"`
}

func ParseRow(data []byte) (*Row, error) {
	var r Row
	if err := yaml.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
