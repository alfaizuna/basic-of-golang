package main

import "fmt"

// TODO:
// 1. Definisikan type Engine struct { Horsepower int }
// 2. Definisikan type Car struct yang meng-embed Engine dan memiliki field Brand (string)

type Engine struct {
	Horsepower int
}

type Car struct {
	Engine
	Brand string
}

func main() {
	// TODO:
	// 1. Buat instance Car:
	//    Brand: "Tesla"
	//    Engine: Engine{Horsepower: 450}
	c := Car{
		Brand: "Tesla",
		Engine: Engine{Horsepower: 450},
	}
	// 2. Cetak promoted field: fmt.Printf("%s berkekuatan %d HP\n", c.Brand, c.Horsepower)
	fmt.Printf("%s berkekuatan %d HP\n", c.Brand, c.Horsepower)
	_ = fmt.Println
}
