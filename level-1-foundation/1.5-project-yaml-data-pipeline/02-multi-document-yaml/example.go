//go:build ignore

package main

import (
	"errors"
	"io"
	"strings"

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

func parseAllRows(yamlContent string) ([]Row, error) {
	decoder := yaml.NewDecoder(strings.NewReader(yamlContent))
	var rows []Row

	for {
		var r Row
		err := decoder.Decode(&r)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		rows = append(rows, r)
	}

	return rows, nil
}
