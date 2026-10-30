package main

import "fmt"

// getLenAndCap mengembalikan panjang (len) dan kapasitas (cap) dari slice s.
func getLenAndCap(s []int) (int, int) {
	// TODO: kembalikan len(s) dan cap(s)
	return len(s), cap(s)
}

func main() {
	data := []int{10, 20, 30, 40, 50}

	// TODO: buat sub-slice 'part' dari data dengan memotong indeks [1:3]
	part := data[1:3]

	// TODO: panggil getLenAndCap(part) dan cetak hasilnya dengan fmt.Println
	l, c := getLenAndCap(part)
	fmt.Println(l, c)
}
