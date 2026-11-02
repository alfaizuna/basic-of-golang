package main

import "fmt"

func countWords(words []string) map[string]int {
	// TODO: hitung frekuensi setiap kata dan kembalikan map hasilnya
	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}
	return freq
}

func main() {
	words := []string{"go", "rust", "go", "python", "go", "rust"}

	// TODO:
	// 1. Panggil counts := countWords(words)
	counts := countWords(words)
	// 2. Cetak: fmt.Println("Go:", counts["go"])
	fmt.Println("Go:", counts["go"])
	// 3. Cetak: fmt.Println("Rust:", counts["rust"])
	fmt.Println("Rust:", counts["rust"])
	// 4. Cetak: fmt.Println("Python:", counts["python"])
	fmt.Println("Python:", counts["python"])

	_ = fmt.Println
	_ = words
}
