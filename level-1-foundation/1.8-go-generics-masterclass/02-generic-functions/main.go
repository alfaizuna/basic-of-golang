package main

import "fmt"

// Last mengembalikan elemen terakhir dari slice T dan boolean status.
func Last[T any](s []T) (T, bool) {
	// TODO: Jika len(s) == 0, kembalikan zero value dan false
	// TODO: Kembalikan elemen terakhir dan true
	var zero T
	return zero, false
}

func main() {
	numbers := []int{10, 20, 30}
	fruits := []string{"apple", "banana", "cherry"}

	lastNum, _ := Last(numbers)
	lastFruit, _ := Last(fruits)

	fmt.Printf("Last number: %d, Last fruit: %s\n", lastNum, lastFruit)
}
