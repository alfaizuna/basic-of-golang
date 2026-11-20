package main

import (
	"encoding/json"
	"fmt"
)

// User merepresentasikan model pengguna dengan format json.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
}

// ToJSON mengonversi User ke string JSON.
func ToJSON(u User) (string, error) {
	_ = json.Marshal
	// TODO: Gunakan json.Marshal
	return "", nil
}

func main() {
	u := User{ID: 1, Username: "gopher"}
	jsonStr, _ := ToJSON(u)
	fmt.Println(jsonStr)
}
