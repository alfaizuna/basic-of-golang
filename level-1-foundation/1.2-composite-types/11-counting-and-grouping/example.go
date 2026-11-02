//go:build ignore

package main

import "fmt"

func main() {
	fruits := []string{"apel", "jeruk", "apel", "mangga", "apel", "jeruk"}

	// Menghitung frekuensi dengan map
	// Saat key belum ada di map, Go mengembalikan 0 (zero value), lalu ditambah 1
	freq := make(map[string]int)
	for _, fruit := range fruits {
		freq[fruit]++
	}

	fmt.Println("Hasil penghitungan:")
	for fruit, count := range freq {
		fmt.Printf("- %s: %d\n", fruit, count)
	}
}
