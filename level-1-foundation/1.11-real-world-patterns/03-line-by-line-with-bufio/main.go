package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// FilterLines membaca baris demi baris dan mengembalikan baris yang memuat keyword.
func FilterLines(r io.Reader, keyword string) ([]string, error) {
	_ = bufio.NewScanner
	_ = strings.Contains
	// TODO: Gunakan bufio.NewScanner dan strings.Contains
	return nil, nil
}

func main() {
	text := "ERROR: connection lost\nINFO: user logged in\nERROR: timeout"
	lines, _ := FilterLines(strings.NewReader(text), "ERROR")
	fmt.Printf("Matched lines count: %d\n", len(lines))
}
