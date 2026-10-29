package main

import "fmt"

// modifyCopy menerima salinan array [3]int, mengubah elemen indeks 0 menjadi 99,
// dan mengembalikan array salinan tersebut.
func modifyCopy(arr [3]int) [3]int {
	// TODO: ubah arr[0] menjadi 99 dan kembalikan arr
	arr[0] = 99
	return arr
}

func main() {
	// 1. Array [3]int berukuran 3
	scores := [3]int{10, 20, 30}

	// 2. Panggil modifyCopy(scores) yang menerima salinan scores
	modified := modifyCopy(scores)

	// 3. Cetak array asli dan array hasil modifikasi
	fmt.Println(scores)
	fmt.Println(modified)
}
