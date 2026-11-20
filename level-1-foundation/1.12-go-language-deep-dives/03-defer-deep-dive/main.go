package main

import "fmt"

// AddTax menambahkan pajak 10% dengan memodifikasi named return di dalam defer.
func AddTax(amount float64) (total float64) {
	// TODO: Inisialisasi total = amount dan modifikasi via defer func()
	return 0.0
}

func main() {
	result := AddTax(100.0)
	fmt.Printf("Total with tax: %.2f\n", result)
}
