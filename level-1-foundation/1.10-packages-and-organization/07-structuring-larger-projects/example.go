//go:build ignore

package main

import "fmt"

func main() {
	fmt.Println("Layout standar proyek Go:")
	fmt.Println("  cmd/          -> entry point biner (misal cmd/api/main.go)")
	fmt.Println("  internal/     -> kode privat yang dilarang diimpor modul lain")
	fmt.Println("  pkg/          -> library publik yang aman dipakai modul luar")
}
