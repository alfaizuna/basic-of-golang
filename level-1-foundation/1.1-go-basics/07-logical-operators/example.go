//go:build ignore

package main

import "fmt"

func main() {
	isWeekend := true
	isRaining := false
	hasUmbrella := true

	canPlayOutside := !isRaining && isWeekend
	canGoOut := !isRaining || hasUmbrella

	fmt.Println(canPlayOutside)
	fmt.Println(canGoOut)
	fmt.Println(!isRaining)
}
