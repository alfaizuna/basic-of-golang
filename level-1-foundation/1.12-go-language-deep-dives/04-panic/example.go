//go:build ignore

package main

import "fmt"

func assertNonEmpty(s string) {
	if s == "" {
		panic("string tidak boleh kosong!")
	}
}

func main() {
	fmt.Println("Sebelum panic")
	assertNonEmpty("ok")
	// assertNonEmpty("") // Menyebabkan crash runtime
}
