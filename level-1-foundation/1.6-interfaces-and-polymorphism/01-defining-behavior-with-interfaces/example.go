//go:build ignore

package main

// example.go — hanya untuk referensi, jangan diubah

type Greeter interface {
	Greet() string
}

type User struct {
	Name string
}

func (u User) Greet() string {
	return "Halo, saya " + u.Name
}

func Welcome(g Greeter) string {
	return g.Greet()
}
