package main

import "fmt"

// example.go — hanya untuk referensi, jangan diubah

type Counter struct {
	Value int
}

func (c Counter) ShowValue() {
	fmt.Println("Value:", c.Value)
}

func (c *Counter) Increment() {
	c.Value++
}

// Contoh penggunaan pointer receiver:
//
//	c := Counter{Value: 0}
//	c.Increment()
//	c.Increment()
//	c.ShowValue() // Value: 2
var _ = fmt.Sprintf // pastikan import fmt tidak unused
