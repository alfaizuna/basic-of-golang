//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Klien dengan konfigurasi timeout 5 detik
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	payload := []byte(`{"title": "Go Concurrency"}`)
	resp, err := client.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Status:", resp.Status)
}
