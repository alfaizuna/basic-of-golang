//go:build ignore

package main

import "fmt"

func main() {
	visited := make(map[string]bool)
	visited["https://example.com"] = true
	fmt.Println("Sudah dikunjungi:", visited["https://example.com"])
}
