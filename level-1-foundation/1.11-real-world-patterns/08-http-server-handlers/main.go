package main

import (
	"fmt"
	"net/http"
)

// PingHandler membalas dengan status 200 dan teks "pong".
func PingHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Set status code OK dan tulis "pong" ke ResponseWriter
}

func main() {
	http.HandleFunc("/ping", PingHandler)
	fmt.Println("Ping handler configured")
}
