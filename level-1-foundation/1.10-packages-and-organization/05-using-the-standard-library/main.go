package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TransformInput memproses string menggunakan paket strings dan strconv bawaan.
func TransformInput(raw string) (int, string, error) {
	_ = strconv.Atoi
	_ = strings.TrimSpace
	// TODO: Gunakan strings.TrimSpace, strings.Fields, dan strconv.Atoi
	return 0, "", nil
}

func main() {
	n, upper, err := TransformInput("  250 usd  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Num: %d, Upper: %s\n", n, upper)
}
