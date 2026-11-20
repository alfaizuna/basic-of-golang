package main

import "fmt"

// TODO: Buat struct generik Pair[T, U any] dengan field First T dan Second U
type Pair[T, U any] struct {
	First  T
	Second U
}

// TODO: Buat fungsi helper NewPair[T, U any](first T, second U) Pair[T, U]
func NewPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{}
}

func main() {
	p := NewPair("HTTP", 200)
	fmt.Printf("Pair: %s = %d\n", p.First, p.Second)
}
