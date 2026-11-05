//go:build ignore

package main

import "fmt"

func double(n int) {
	n = n * 2
	fmt.Println("Di dalam fungsi double:", n) // 20
}

func main() {
	x := 10
	double(x)
	fmt.Println("Di dalam main:", x) // tetap 10!
}
