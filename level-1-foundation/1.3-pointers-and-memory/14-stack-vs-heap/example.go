package main

// example.go — hanya untuk referensi, jangan diubah

type Point struct {
	X, Y int
}

func staysOnStack() int {
	p := Point{X: 10, Y: 20}
	return p.X + p.Y
}

func escapesToHeap() *Point {
	p := Point{X: 10, Y: 20}
	return &p
}
