package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

// SendJSONPayload mengirimkan byte JSON ke URL target menggunakan http.Client.
func SendJSONPayload(client *http.Client, url string, jsonData []byte) (int, error) {
	_ = bytes.NewBuffer
	// TODO: Lakukan client.Post, defer resp.Body.Close(), dan kembalikan resp.StatusCode
	return 0, nil
}

func main() {
	client := &http.Client{Timeout: 3 * time.Second}
	status, err := SendJSONPayload(client, "https://example.com", []byte(`{"hello":"world"}`))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Created status: %d\n", status)
}
