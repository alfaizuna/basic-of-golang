//go:build ignore

package main

import "fmt"

func main() {
	// Buffered channel dengan kapasitas 2
	ch := make(chan string, 2)

	// Pengiriman tidak memblokir karena kapasitas buffer masih muat
	ch <- "pesan 1"
	ch <- "pesan 2"

	fmt.Println("Panjang buffer saat ini:", len(ch))
	fmt.Println("Kapasitas total channel:", cap(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
