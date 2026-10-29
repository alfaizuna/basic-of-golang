//go:build ignore

package main

import "fmt"

const Pi = 3.14

const (
	A = iota // 0
	B        // 1
	C        // 2
)

func main() {
	var count int   // zero value: 0
	var isDone bool // zero value: false

	fmt.Println(Pi)
	fmt.Println(A, B, C)
	fmt.Println(count, isDone)
}
