package main

import "fmt"

// TODO: Buat interface constraint Measurable dengan method Area() float64
type Measurable interface {
}

type Square struct {
	Side float64
}

// TODO: Implementasikan method Area() float64 pada Square

// SumAreas menjumlahkan nilai Area dari sekumpulan objek Measurable.
func SumAreas[T Measurable](items []T) float64 {
	// TODO: Lakukan perulangan dan jumlahkan item.Area()
	return 0.0
}

func main() {
	squares := []Square{{Side: 3}, {Side: 4}}
	total := SumAreas(squares)
	fmt.Printf("Total Area: %.2f\n", total)
}
