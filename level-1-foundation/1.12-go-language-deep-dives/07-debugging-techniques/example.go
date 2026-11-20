//go:build ignore

package main

import "fmt"

func buggyFunction(items []int) int {
	// Bug: index out of range jika slice kosong
	return items[0]
}

func main() {
	fmt.Println("Pemeriksaan alat bantu: go vet ./...")
}
