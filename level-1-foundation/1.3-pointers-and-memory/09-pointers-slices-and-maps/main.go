package main

import "fmt"

// addBonus seharusnya menambahkan bonus ke tiap elemen slice.
// Tapi ada bug: v adalah salinan, bukan elemen asli!
func addBonus(scores []int, bonus int) {
	for i := range scores {
		scores[i] += bonus
	}
}

// resetMap seharusnya mengosongkan semua isi map.
// Tapi ada bug: assign map baru tidak mengubah map yang dikirim pemanggil!
func resetMap(m map[string]int) {
	for k := range m {
		delete(m, k)
	}
}

func main() {
	scores := []int{10, 20, 30}
	addBonus(scores, 5)
	fmt.Println("scores:", scores)

	stock := map[string]int{"apel": 10, "jeruk": 5}
	resetMap(stock)
	fmt.Println("stock:", stock)
}
