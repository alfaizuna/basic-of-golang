//go:build ignore

package main

import "fmt"

// Constraint interface dengan method
type Stringer interface {
	String() string
}

func PrintAllStringers[T Stringer](items []T) {
	for _, item := range items {
		fmt.Println(item.String())
	}
}

type User struct {
	Name string
}

func (u User) String() string {
	return "User: " + u.Name
}

func main() {
	users := []User{{Name: "Alice"}, {Name: "Bob"}}
	PrintAllStringers(users)
}
