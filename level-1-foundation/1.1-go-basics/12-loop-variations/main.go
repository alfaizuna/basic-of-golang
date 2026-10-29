package main

import "fmt"

// 1. Buat fungsi countVowels(text string) int
// Gunakan for _, char := range text:
// - Cek apakah char adalah 'a', 'i', 'u', 'e', atau 'o'
// - Hitung dan kembalikan total huruf vokal
func countVowels(text string) int {
	total := 0
	for _, char := range text {
		switch char {
		case 'a', 'i', 'u', 'e','o':
			total++
		}
	}
	return total
}

// 2. Buat fungsi findTarget(matrixSize int, targetSum int) (int, int)
// Gunakan loop bersarang dengan label di loop terluar:
// - r dari 1 s/d matrixSize, c dari 1 s/d matrixSize
// - Jika r + c == targetSum, simpan nilainya dan break label luar
// - Kembalikan foundR dan foundC

func findTarget(matrixSize int, targetSum int) (int, int) {
	foundR, foundC := 0, 0
	SearchLoop:
	for r := 1; r <= matrixSize; r++ {
		for c := 1; c <= matrixSize; c++ {
			if r+c == targetSum {
				foundR = r
				foundC = c
				break SearchLoop
			}
		}
	}
	return foundR, foundC
}

func main() {
	// 3. Panggil countVowels("golang") dan cetak hasilnya
	fmt.Println(countVowels("golang"))

	// 4. Panggil findTarget(5, 7) dan cetak hasilnya (r dan c)
	r, c := findTarget(5, 7)
	fmt.Println(r, c)
}
