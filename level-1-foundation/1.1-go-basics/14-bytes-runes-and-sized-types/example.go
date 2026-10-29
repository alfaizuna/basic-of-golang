//go:build ignore

package main

import "fmt"

func main() {
	// 1. byte adalah alias untuk uint8 (0 s/d 255)
	var b byte = 'A' // kode ASCII 65
	fmt.Printf("byte: %d (%c)\n", b, b)

	// 2. rune adalah alias untuk int32 (Unicode Code Point)
	var r rune = '🚀' // kode Unicode 128640
	fmt.Printf("rune: %d (%c)\n", r, r)

	// 3. Sized integer & overflow wrap-around
	var maxInt8 int8 = 127
	var nextInt8 int8 = maxInt8 + 1 // overflow: berputar ke nilai minimum (-128)
	fmt.Printf("int8: %d + 1 = %d\n", maxInt8, nextInt8)
}
