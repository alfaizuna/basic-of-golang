package main

import "fmt"

func PrintAll(items []any) {
	for _, item := range items {
		fmt.Printf("[%T] %v\n", item, item)
	}
}

func main() {
	data := []any{100, "Go is awesome", true}
	PrintAll(data)
}
