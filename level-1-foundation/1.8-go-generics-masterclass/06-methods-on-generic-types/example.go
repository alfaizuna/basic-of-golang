//go:build ignore

package main

import "fmt"

type Box[T any] struct {
	value T
}

// Method receiver menyertakan parameter tipe [T]
func (b *Box[T]) Set(val T) {
	b.value = val
}

func (b *Box[T]) Get() T {
	return b.value
}

func main() {
	intBox := Box[int]{}
	intBox.Set(42)
	fmt.Println("Isi box:", intBox.Get())
}
