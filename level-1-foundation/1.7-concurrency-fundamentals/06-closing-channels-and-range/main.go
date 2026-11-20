package main

import "fmt"

// GenerateFibonacci menghasilkan n angka Fibonacci pertama melalui channel tertutup.
func GenerateFibonacci(n int) <-chan int {
	// TODO: Buat channel int berkapasitas n
	// TODO: Luncurkan goroutine yang mengirim deret Fibonacci dan menutup channel
	return nil
}

func main() {
	var results []int
	for num := range GenerateFibonacci(5) {
		results = append(results, num)
	}
	fmt.Printf("Fibonacci: %v\n", results)
}
