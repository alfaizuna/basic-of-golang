//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Halo Pengunjung!")
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	fmt.Println("Server running on :8080...")
	// http.ListenAndServe(":8080", nil)
}
