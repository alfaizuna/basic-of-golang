//go:build ignore

package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type UserProfile struct {
	Username string
	Addr     Address  // Struct di dalam struct (nested)
	Skills   []string // Slice sebagai field
}

func main() {
	profile := UserProfile{
		Username: "gopher123",
		Addr: Address{
			City:    "Jakarta",
			Country: "Indonesia",
		},
		Skills: []string{"Go", "PostgreSQL", "Docker"},
	}

	// Akses field bertingkat
	fmt.Println("Username:", profile.Username)
	fmt.Println("Kota:", profile.Addr.City)
	fmt.Println("Skill utama:", profile.Skills[0])
}
