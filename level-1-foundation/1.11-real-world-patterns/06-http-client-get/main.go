package main

import (
	"fmt"
	"io"
	"net/http"
)

// FetchStatusCodeAndBody melakukan request HTTP GET dan mengembalikan status code beserta body.
func FetchStatusCodeAndBody(url string) (int, string, error) {
	_ = http.Get
	_ = io.ReadAll
	// TODO: Lakukan http.Get, defer resp.Body.Close(), dan io.ReadAll
	return 0, "", nil
}

func main() {
	status, _, err := FetchStatusCodeAndBody("https://example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Status: %d\n", status)
}
