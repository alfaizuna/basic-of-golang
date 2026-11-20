//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Success bool   `json:"success"`
}

func main() {
	raw := `{"status": "ok", "code": 200, "success": true}`
	var resp Response

	err := json.Unmarshal([]byte(raw), &resp) // Wajib passing pointer struct
	if err != nil {
		fmt.Println("Gagal decode:", err)
		return
	}
	fmt.Printf("Parsed: %+v\n", resp)
}
