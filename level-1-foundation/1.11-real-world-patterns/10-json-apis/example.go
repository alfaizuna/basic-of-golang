//go:build ignore

package main

import (
	"encoding/json"
	"net/http"
)

type CalcRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type CalcResponse struct {
	Result int `json:"result"`
}

func main() {
	// Server JSON handler
}
