package main

import "fmt"

// double menerima pointer n dan mengalikan nilainya dengan 2
func double(n *int) {
	// TODO: cek jika n != nil, lalu kalikan nilai yang ditunjuk dengan 2

	if n != nil {
		*n *= 2
	}

	_ = n
}

func main() {
	score := 25

	// TODO: Panggil fungsi double dengan mengirim alamat score (&score)
	double(&score)

	// TODO: Cetak dengan fmt.Println("Score setelah didouble:", score)
	fmt.Println("Score setelah didouble:", score)

	_ = score
	_ = fmt.Println
}
