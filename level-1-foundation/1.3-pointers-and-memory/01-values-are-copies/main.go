package main

import "fmt"

// resetCount menerima int salinan dan mencoba mereset nilainya menjadi 0.
func resetCount(count int) {
	// TODO: ubah count menjadi 0
	// TODO: cetak "Di dalam resetCount:", count
	count = 0
	fmt.Println("Di dalam resetCount:", count)
}

func main() {
	total := 100

	// TODO: panggil resetCount(total)
	resetCount(total)

	// TODO: cetak "Di dalam main:", total
	fmt.Println("Di dalam main:", total)

	_ = total
	_ = fmt.Println
}
