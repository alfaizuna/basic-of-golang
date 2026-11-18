package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
	Role string
}

var userDatabase = []User{
	{ID: 1, Name: "Alice", Role: "Admin"},
	{ID: 2, Name: "Bob", Role: "Member"},
}

// FindUser mencari user berdasarkan ID di userDatabase.
// Jika user ditemukan, kembalikan pointer ke User tersebut dan nil error.
// Jika user tidak ditemukan, kembalikan nil dan error dengan pesan "user not found".
func FindUser(id int) (*User, error) {
	for i := range userDatabase {
		if userDatabase[i].ID == id {
			return &userDatabase[i], nil
		}
	}
	// TODO: Kembalikan nil dan error dengan pesan "user not found"
	return nil, errors.New("user not found")
}

// FormatUserName mengembalikan string nama pengguna.
// Jika u adalah nil, kembalikan "Guest".
// Jika u tidak nil, kembalikan u.Name.
func FormatUserName(u *User) string {
	// TODO: Tambahkan nil guard jika u == nil, kembalikan "Guest"
	if u == nil {
		return "Guest"
	}
	return u.Name
}

func main() {
	// 1. Cari user yang ada (ID 1)
	user1, err := FindUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User:", FormatUserName(user1))
	}

	// 2. Cari user yang tidak ada (ID 99)
	user99, err := FindUser(99)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User:", FormatUserName(user99))
	}

	// 3. Panggil FormatUserName dengan pointer nil secara eksplisit
	fmt.Println("Fallback:", FormatUserName(nil))
}
