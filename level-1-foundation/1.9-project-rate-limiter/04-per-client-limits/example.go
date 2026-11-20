//go:build ignore

package main

import "fmt"

type ClientState struct {
	count int
}

func main() {
	clients := make(map[string]*ClientState)
	clients["user-123"] = &ClientState{count: 1}
	fmt.Println("Client count:", clients["user-123"].count)
}
