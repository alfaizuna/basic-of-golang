package main

import "fmt"

// createBufferedSlice membuat slice integer dengan panjang dan kapasitas tertentu.
func createBufferedSlice(length int, capacity int) []int {
	// TODO: gunakan make([]int, length, capacity)
	return make([]int, length, capacity)
}

func main() {
	// TODO:
	// 1. Buat slice 's' dengan memanggil createBufferedSlice(2, 5)
	s := createBufferedSlice(2, 5)
	// 2. Isi s[0] = 100 dan s[1] = 200
	s[0] = 100
	s[1] = 200
	// 3. Cetak slice: fmt.Println(s)
	fmt.Println(s)
	// 4. Cetak panjang dan kapasitas: fmt.Println(len(s), cap(s))
	fmt.Println(len(s), cap(s))
}
