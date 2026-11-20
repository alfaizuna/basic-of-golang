package main

import "fmt"

// Stack adalah struktur data tumpukan Last-In-First-Out berpengetikan generik.
type Stack[T any] struct {
	items []T
}

// Push menambahkan elemen ke tumpukan.
func (s *Stack[T]) Push(val T) {
	// TODO: Tambahkan val ke slice s.items
}

// Pop mengambil elemen teratas dari tumpukan.
func (s *Stack[T]) Pop() (T, bool) {
	// TODO: Jika kosong kembalikan zero value dan false, jika ada kembalikan elemen dan true
	var zero T
	return zero, false
}

// Len mengembalikan jumlah elemen di tumpukan.
func (s *Stack[T]) Len() int {
	// TODO: Kembalikan jumlah elemen s.items
	return 0
}

func main() {
	var stack Stack[int]
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	val, _ := stack.Pop()
	fmt.Printf("Pop: %d, Len: %d\n", val, stack.Len())
}
