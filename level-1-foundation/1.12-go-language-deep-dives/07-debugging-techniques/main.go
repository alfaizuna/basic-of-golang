package main

import "fmt"

// SafeFirst mengembalikan elemen pertama secara aman tanpa risiko runtime panic.
func SafeFirst(items []int) (int, bool) {
	// TODO: Tangani slice kosong secara aman
	return 0, false
}

func main() {
	val, ok := SafeFirst([]int{10, 20})
	fmt.Printf("First: %d, OK: %t\n", val, ok)
}
