package main

import "fmt"

func addScore(scores []int, newScore int) []int {
	// TODO: tambahkan newScore ke scores menggunakan append dan kembalikan hasilnya
	return append(scores, newScore)
}

func main() {
	// TODO:
	// 1. Buat slice scores ([]int) berisi nilai awal 80 dan 90
	scores := []int{80, 90}
	// 2. Tambahkan 100 ke scores menggunakan addScore
	scores = addScore(scores, 100)
	// 3. Tambahkan 95 ke scores menggunakan addScore
	scores = addScore(scores,95)
	// 4. Cetak: fmt.Println("Scores:", scores)
	fmt.Println("Scores:", scores)
	// 5. Cetak: fmt.Println("Length:", len(scores))
	fmt.Println("Length:", len(scores))

	_ = fmt.Println
}
