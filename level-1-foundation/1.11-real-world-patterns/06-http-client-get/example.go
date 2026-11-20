//go:build ignore

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("Request gagal:", err)
		return
	}
	defer resp.Body.Close() // Wajib ditutup!

	fmt.Println("Status Code:", resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	_ = body
}
