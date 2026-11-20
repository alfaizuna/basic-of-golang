//go:build ignore

package main

import "fmt"

func main() {
	// Membuat unbuffered channel bertipe string
	messages := make(chan string)

	// Goroutine pengirim (sender)
	go func() {
		messages <- "ping" // Mengirim data ke channel
	}()

	// Goroutine penerima (receiver di goroutine main)
	msg := <-messages // Menerima data dari channel (blocking hingga pengirim siap)
	fmt.Println("Diterima:", msg)
}
