package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// 1. countBytesAndRunes menghitung jumlah byte dan jumlah rune pada text.
func countBytesAndRunes(text string) (int, int) {
	byteCount := len(text)
	// TODO: gunakan utf8.RuneCountInString(text) untuk menghitung rune sebenarnya
	runeCount := utf8.RuneCountInString(text) // ganti "" dengan text
	return byteCount, runeCount
}

// 2. normalizeText merapikan spasi di ujung teks dan mengubahnya ke huruf kecil.
func normalizeText(text string) string {
	trimmed := strings.TrimSpace(text)
	// TODO: ubah trimmed menjadi huruf kecil dengan strings.ToLower
	return strings.ToLower(trimmed)
}

func main() {
	// Contoh pemanggilan 1: menghitung byte dan rune "Go 🚀"
	b, r := countBytesAndRunes("Go 🚀")
	fmt.Println(b, r)

	// TODO: Panggil normalizeText("  Belajar GOLANG  ") dan cetak hasilnya
	fmt.Println(normalizeText("  Belajar GOLANG  "))
}
