package main

import "fmt"

func SumValues(a, b int) int {
	return a + b
}

func CreateBuffer(size int) *[]byte {
	buf := make([]byte, size)
	return &buf
}

func main() {
	s := SumValues(10, 20)
	fmt.Println("Sum:", s)

	buf := CreateBuffer(1024)
	if buf != nil {
		fmt.Println("Buffer len:", len(*buf))
	}
}
