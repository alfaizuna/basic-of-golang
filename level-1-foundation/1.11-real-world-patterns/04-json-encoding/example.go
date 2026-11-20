//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags,omitempty"`
}

func main() {
	p := Product{ID: 1, Name: "Keyboard", Price: 75.5}
	bytes, _ := json.Marshal(p)
	fmt.Println(string(bytes))
}
