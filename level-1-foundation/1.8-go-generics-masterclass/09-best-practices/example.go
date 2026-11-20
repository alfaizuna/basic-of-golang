//go:build ignore

package main

import "fmt"

// BAIK: Generics pada struktur data koleksi universal
type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

func (s *Set[T]) Add(val T) {
	s.elements[val] = struct{}{}
}

func main() {
	s := NewSet[string]()
	s.Add("Go")
	fmt.Println("Set berhasil dibuat")
}
