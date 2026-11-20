package main

import "fmt"

// TODO: Definisikan interface OrderedNumber dengan ~int | ~float64
type OrderedNumber interface {
}

type Speed float64

// Max mengembalikan nilai terbesar dari dua nilai bertipe OrderedNumber.
func Max[T OrderedNumber](a, b T) T {
	// TODO: Kembalikan nilai terbesar antara a dan b
	return a
}

func main() {
	s1 := Speed(60.0)
	s2 := Speed(85.5)
	maxSpeed := Max(s1, s2)
	fmt.Printf("Max speed: %.2f\n", maxSpeed)
}
