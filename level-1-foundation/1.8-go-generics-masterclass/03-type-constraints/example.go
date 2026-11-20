//go:build ignore

package main

import "fmt"

// comparable mengizinkan tipe apapun yang mendukung perbandingan == dan !=
func IndexOf[T comparable](s []T, target T) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	langs := []string{"Go", "Python", "Rust"}
	fmt.Println("Index of Rust:", IndexOf(langs, "Rust"))
}
