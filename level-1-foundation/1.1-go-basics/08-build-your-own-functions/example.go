//go:build ignore

package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sayHello(name string) {
	fmt.Println("Halo,", name)
}

func main() {
	total := add(10, 5)
	fmt.Println("Total:", total)

	sayHello("Gopher")
}
