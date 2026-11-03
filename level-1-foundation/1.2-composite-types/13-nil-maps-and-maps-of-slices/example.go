//go:build ignore

package main

import "fmt"

func main() {
	// 1. Nil map
	var nilMap map[string]int
	// Membaca nil map aman: mengembalikan zero value (0)
	fmt.Println("Membaca nil map:", nilMap["tidak_ada"])

	// PERINGATAN: nilMap["kunci"] = 10 -> PANIC: assignment to entry in nil map!
	// Selalu inisialisasi dengan make() sebelum menulis:
	safeMap := make(map[string]int)
	safeMap["kunci"] = 10
	fmt.Println("safeMap:", safeMap)

	// 2. Map of Slices: map[string][]string
	teams := make(map[string][]string)
	teams["Backend"] = append(teams["Backend"], "Budi", "Siti")
	teams["Frontend"] = append(teams["Frontend"], "Andi")

	fmt.Println("Backend team:", teams["Backend"])
	fmt.Println("Frontend team:", teams["Frontend"])
	fmt.Println(teams)
}
