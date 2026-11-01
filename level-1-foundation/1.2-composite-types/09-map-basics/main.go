package main

import "fmt"

func main() {
	// TODO:
	// 1. Buat map scores := map[string]int{"Alice": 90, "Bob": 85}
	scores := map[string]int{"Alice": 90, "Bob": 85}
	// 2. Tambahkan Charlie dengan nilai 95
	scores["Charlie"] = 95
	// 3. Hapus Bob menggunakan delete
	delete(scores, "Bob")
	// 4. Cetak: fmt.Println("Alice:", scores["Alice"])
	fmt.Println("Alice:", scores["Alice"])
	// 5. Cetak: fmt.Println("Total:", len(scores))
	fmt.Println("Total:", len(scores))

	_ = fmt.Println
}
