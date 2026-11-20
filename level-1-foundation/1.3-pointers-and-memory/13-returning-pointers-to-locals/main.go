package main

import "fmt"

type UserProfile struct {
	Username string
	Email    string
}

// NewCounter mengembalikan pointer ke nilai integer lokal.
// Di Go, mengembalikan alamat variabel lokal aman karena akan escape ke heap.
func NewCounter(initial int) *int {
	// TODO: deklarasikan variabel lokal int dengan nilai initial
	// lalu kembalikan pointernya
	return nil
}

// NewProfile mengembalikan pointer ke struct UserProfile lokal.
func NewProfile(username, email string) *UserProfile {
	// TODO: buat struct UserProfile dengan field Username dan Email,
	// lalu kembalikan pointernya
	return nil
}

func main() {
	c := NewCounter(10)
	if c != nil {
		fmt.Println("Counter:", *c)
	}

	p := NewProfile("budi", "budi@example.com")
	if p != nil {
		fmt.Printf("Profile: %s (%s)\n", p.Username, p.Email)
	}
}
