//go:build ignore

package main

import "fmt"

func lifoExample() {
	// Eksekusi Last-In, First-Out (tumpukan terbalik)
	defer fmt.Println("1 (pertama didaftarkan, terakhir dieksekusi)")
	defer fmt.Println("2")
	defer fmt.Println("3 (terakhir didaftarkan, pertama dieksekusi)")
}

func main() {
	lifoExample()
}
