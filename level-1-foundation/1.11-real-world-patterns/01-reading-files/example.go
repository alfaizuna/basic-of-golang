//go:build ignore

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("example.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File tidak ditemukan!")
		return
	} else if err != nil {
		fmt.Println("Gagal membaca file:", err)
		return
	}
	fmt.Println("Isi file:", string(data))
}
