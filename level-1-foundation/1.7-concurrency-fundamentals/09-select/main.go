package main

import "fmt"

// FirstResponder mengembalikan pesan pertama yang tiba dari dua channel.
func FirstResponder(chA, chB <-chan string) string {
	// TODO: Gunakan select untuk mendengarkan chA dan chB, lalu kembalikan pesan yang pertama tiba
	return ""
}

func main() {
	chA := make(chan string, 1)
	chB := make(chan string, 1)

	chA <- "Fast response"
	result := FirstResponder(chA, chB)
	fmt.Printf("First arrived: %s\n", result)
}
