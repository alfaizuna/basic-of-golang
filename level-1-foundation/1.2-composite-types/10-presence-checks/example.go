//go:build ignore

package main

import "fmt"

func main() {
	scores := map[string]int{
		"Alice": 100,
		"Bob":   0, // Bob nilainya 0 (zero value)
	}

	// Masalah jika hanya membaca m[key]:
	// scores["Bob"] = 0
	// scores["Charlie"] = 0 (karena Charlie tidak ada, Go mengembalikan zero value int)
	fmt.Println("scores['Bob']:", scores["Bob"])
	fmt.Println("scores['Charlie']:", scores["Charlie"])

	// Solusi: Pola comma-ok idiom
	val, ok := scores["Charlie"]
	if ok {
		fmt.Println("Charlie ada dengan nilai:", val)
	} else {
		fmt.Println("Charlie tidak ada di map!")
	}
}
