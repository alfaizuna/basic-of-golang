//go:build ignore

package main

import "fmt"

func describe(val any) {
	fmt.Printf("(%v, %T)\n", val, val)
}

func main() {
	var i any = 42
	describe(i)

	i = "halo gopher"
	describe(i)

	i = struct {
		Name string
	}{Name: "Gopher"}
	describe(i)
}
