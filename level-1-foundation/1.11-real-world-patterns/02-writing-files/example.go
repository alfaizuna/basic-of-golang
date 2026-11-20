//go:build ignore

package main

import (
	"fmt"
	"os"
)

func main() {
	// Menulis data langsung
	err := os.WriteFile("output.txt", []byte("Halo dari Go!"), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File berhasil ditulis")
}
