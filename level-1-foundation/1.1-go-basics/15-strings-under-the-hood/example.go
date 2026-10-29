//go:build ignore

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// 1. String adalah immutable slice of bytes (UTF-8 encoded)
	greeting := "Halo 世界"

	// len() mengembalikan jumlah BYTE, bukan jumlah karakter
	fmt.Printf("Jumlah byte: %d\n", len(greeting)) // 11 byte ('世' dan '界' masing-masing 3 byte)

	// utf8.RuneCountInString() menghitung jumlah karakter (rune) sebenarnya
	fmt.Printf("Jumlah rune: %d\n", utf8.RuneCountInString(greeting)) // 7 rune

	// 2. Iterasi rune:
	// for i, r := range s otomatis men-decode UTF-8 menjadi rune
	for idx, char := range greeting {
		fmt.Printf("Byte index: %d, Char: %c\n", idx, char)
	}

	// 3. Paket strings untuk manipulasi teks
	fmt.Println(strings.ToLower("GO LANGUAGE"))
	fmt.Println(strings.TrimSpace("  bersih  "))
}
