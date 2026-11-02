//go:build ignore

package main

import (
	"fmt"
	"sort"
)

func main() {
	scores := map[string]int{
		"Charlie": 85,
		"Alice":   95,
		"Bob":     90,
	}

	// 1. Iterasi langsung pada map menghasilkan urutan acak!
	fmt.Println("--- Iterasi Langsung (Acak) ---")
	for k, v := range scores {
		fmt.Printf("%s: %d\n", k, v)
	}

	// 2. Iterasi Terurut: Ambil semua key, urutkan, lalu loop key-nya
	var keys []string
	for k := range scores {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("\n--- Iterasi Terurut ---")
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, scores[k])
	}
}
