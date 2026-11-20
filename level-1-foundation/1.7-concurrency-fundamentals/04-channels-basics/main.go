package main

import "fmt"

// ComputeSquare menghitung kuadrat n menggunakan goroutine dan channel unbuffered.
func ComputeSquare(n int) int {
	// TODO: Buat unbuffered channel int
	// TODO: Luncurkan goroutine untuk mengirim n * n ke channel
	// TODO: Terima dan kembalikan nilai dari channel
	return 0
}

func main() {
	val := 9
	result := ComputeSquare(val)
	fmt.Printf("Kuadrat dari %d adalah %d\n", val, result)
}
