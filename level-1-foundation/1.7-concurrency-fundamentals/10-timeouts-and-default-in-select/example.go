//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	// Non-blocking receive dengan default
	select {
	case msg := <-ch:
		fmt.Println("Menerima:", msg)
	default:
		fmt.Println("Channel kosong, lanjutkan tanpa memblokir!")
	}

	// Operasi dengan timeout
	select {
	case msg := <-ch:
		fmt.Println("Menerima:", msg)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("Timeout! Operasi dibatalkan setelah 50ms")
	}
}
