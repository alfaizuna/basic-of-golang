//go:build ignore

package main

import "fmt"

func main() {
	temperature := 30
	if temperature > 35 {
		fmt.Println("Panas sekali")
	} else if temperature > 25 {
		fmt.Println("Hangat nyaman")
	} else {
		fmt.Println("Sejuk")
	}

	if status := 200; status == 200 {
		fmt.Println("OK")
	}
}
