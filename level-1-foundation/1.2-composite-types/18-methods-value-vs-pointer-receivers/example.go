//go:build ignore

package main

import "fmt"

type Counter struct {
	Count int
}

// 1. Value Receiver: menerima salinan (copy).
// Perubahan di dalam method TIDAK mengubah struct asli.
func (c Counter) Display() {
	fmt.Println("Count sekarang:", c.Count)
}

// 2. Pointer Receiver: menerima alamat memori (*Counter).
// Perubahan di dalam method LANGSUNG mengubah struct asli!
func (c *Counter) Increment() {
	c.Count++
}

func main() {
	c := Counter{Count: 0}

	c.Increment() // Otomatis (&c).Increment() oleh Go
	c.Increment()
	c.Display() // Output: 2
}
