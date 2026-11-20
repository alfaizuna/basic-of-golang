package main

// example.go — hanya untuk referensi, jangan diubah

type Counter struct {
	Count int
}

func (c *Counter) Increment() {
	if c == nil {
		return
	}
	c.Count++
}

func updateCount(c *Counter, amount int) {
	if c == nil {
		return
	}
	c.Count += amount
}
