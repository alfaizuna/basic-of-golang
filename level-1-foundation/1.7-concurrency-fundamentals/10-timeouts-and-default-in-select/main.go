package main

import (
	"fmt"
	"time"
)

// FetchWithTimeout mengambil pesan dari channel dengan batas waktu tertentu.
func FetchWithTimeout(ch <-chan string, timeout time.Duration) (string, error) {
	// TODO: Gunakan select dengan case res := <-ch dan case <-time.After(timeout)
	return "", nil
}

func main() {
	ch := make(chan string, 1)
	ch <- "Success!"

	res, err := FetchWithTimeout(ch, 100*time.Millisecond)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Result: %s\n", res)
}
