//go:build ignore

package main

import "fmt"

// Deklarasi tipe struct
type User struct {
	Name string
	Age  int
}

func main() {
	// 1. Inisialisasi dengan named fields (rekomendasi idiom Go)
	u1 := User{
		Name: "Alice",
		Age:  28,
	}

	// 2. Akses dan modifikasi field
	fmt.Printf("User: %s, Umur: %d\n", u1.Name, u1.Age)

	u1.Age = 29
	fmt.Println("Setelah ulang tahun:", u1.Age)

	// 3. Zero value struct (semua field bernilai zero value masing-masing)
	var u2 User
	fmt.Printf("Zero value User: %+v\n", u2) // Name: "", Age: 0
}
