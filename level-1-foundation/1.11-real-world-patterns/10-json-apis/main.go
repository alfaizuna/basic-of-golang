package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	Sum int `json:"sum"`
}

// SumHandler menerima request JSON dan mengembalikan penjumlahan nilainya dalam format JSON.
func SumHandler(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder
	_ = json.NewEncoder
	// TODO: Set header Content-Type, decode request, hitung jumlah, dan encode response
}

func main() {
	http.HandleFunc("/api/sum", SumHandler)
	fmt.Println("API sum endpoint ready")
}
