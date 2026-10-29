//go:build ignore

package main

import "fmt"

func main() {
	intDiv := 7 / 2
	mod := 7 % 2
	floatDiv := 7.0 / 2.0

	var a int = 10
	var b float64 = 2.5
	total := float64(a) + b

	fmt.Println(intDiv)
	fmt.Println(mod)
	fmt.Println(floatDiv)
	fmt.Println(total)
}
