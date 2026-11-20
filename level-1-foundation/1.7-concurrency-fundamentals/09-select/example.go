//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- "dari satu"
	}()

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- "dari dua"
	}()

	// Menunggu mana yang siap lebih dulu
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Menerima:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Menerima:", msg2)
		}
	}
}
