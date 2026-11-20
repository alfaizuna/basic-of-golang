//go:build ignore

package main

import "fmt"

// Constraint yang menerima tipe bertipe dasar int atau float64
type Number interface {
	~int | ~float64
}

type MyInt int // Tipe kustom dengan underlying type int

func Double[T Number](n T) T {
	return n * 2
}

func main() {
	var a int = 10
	var b MyInt = 20
	var c float64 = 4.5

	fmt.Println(Double(a))
	fmt.Println(Double(b)) // Valid berkat simbol ~
	fmt.Println(Double(c))
}
