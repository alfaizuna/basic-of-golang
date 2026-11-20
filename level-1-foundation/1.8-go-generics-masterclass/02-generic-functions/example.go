//go:build ignore

package main

import "fmt"

// First mengembalikan elemen pertama dari slice bertipe apapun
func First[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[0], true
}

func main() {
	numbers := []int{10, 20, 30}
	names := []string{"Alice", "Bob"}

	// Type inference: compiler otomatis mengetahui tipe T dari argumen
	n, _ := First(numbers)
	name, _ := First(names)

	fmt.Printf("Angka: %d, Nama: %s\n", n, name)
}
