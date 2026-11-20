package main

import "fmt"

// AccumulateConcurrent menjumlahkan semua elemen secara konkuren tanpa data race.
func AccumulateConcurrent(values []int) int {
	// TODO: Gunakan channel berkapasitas len(values) untuk mengumpulkan nilai dari setiap goroutine
	// TODO: Jumlahkan nilai yang diterima dan kembalikan total
	return 0
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := AccumulateConcurrent(numbers)
	fmt.Printf("Total sum: %d\n", sum)
}
