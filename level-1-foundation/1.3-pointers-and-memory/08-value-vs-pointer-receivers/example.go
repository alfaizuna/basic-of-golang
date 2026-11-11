package main

import "fmt"

// example.go — hanya untuk referensi, jangan diubah

type Rectangle struct {
	Width  float64
	Height float64
}

// Value receiver — hanya membaca
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Value receiver — hanya membaca
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Pointer receiver — mengubah struct
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

var _ = fmt.Sprintf // pastikan import fmt tidak unused
