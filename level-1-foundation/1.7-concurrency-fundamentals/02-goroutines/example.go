//go:build ignore

package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	fmt.Println("Halo dari goroutine:", name)
}

func main() {
	// Menjalankan sayHello secara konkuren
	go sayHello("Gopher")

	// Goroutine fungsi anonim
	go func(msg string) {
		fmt.Println("Anonim goroutine:", msg)
	}("Belajar Go")

	// Beri jeda sejenak agar goroutine sempat berjalan sebelum main exit
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Main selesai")
}
