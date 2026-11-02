package main

import (
	"fmt"
	"sort"
)

func main() {
	scores := map[string]int{"Charlie": 75, "Alice": 95, "Bob": 85}

	// TODO:
	// 1. Kumpulkan semua key dari scores ke dalam slice keys ([]string)
	var keys []string
	for k := range scores {
		keys = append(keys, k)
	}
	// 2. Urutkan keys secara alfabetis dengan sort.Strings(keys)
	sort.Strings(keys)
	// 3. Loop keys yang sudah terurut dan cetak: fmt.Printf("%s: %d\n", k, scores[k])
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, scores[k])
	}

	_ = fmt.Println
	_ = sort.Strings
	_ = scores
}
