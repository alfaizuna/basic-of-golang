//go:build ignore

package main

import "fmt"

type Player struct {
	Name  string
	Level int
}

// levelUp menerima pointer ke struct Player (*Player)
func levelUp(p *Player) {
	if p == nil {
		return
	}
	// Auto-dereference: Go otomatis mengubah p.Level menjadi (*p).Level
	p.Level++
}

func main() {
	hero := Player{Name: "Arthur", Level: 1}
	fmt.Println("Sebelum level up:", hero) // {Arthur 1}

	// Mengirimkan pointer ke struct hero (&hero)
	levelUp(&hero)

	fmt.Println("Setelah level up:", hero) // {Arthur 2}
}
