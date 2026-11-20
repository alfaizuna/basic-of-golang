//go:build ignore

package main

import "fmt"

func safeCall() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Berhasil pulih dari panic:", r)
		}
	}()

	fmt.Println("Melakukan operasi berbahaya...")
	panic("koneksi database meledak")
	fmt.Println("Baris ini tidak akan pernah tercapai")
}

func main() {
	safeCall()
	fmt.Println("Program utama tetap hidup dan berjalan normal!")
}
