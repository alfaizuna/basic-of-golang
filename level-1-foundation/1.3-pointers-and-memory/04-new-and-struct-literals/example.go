//go:build ignore

package main

import "fmt"

type Database struct {
	Name    string
	Timeout int
}

func main() {
	// 1. Menggunakan new() untuk tipe dasar
	// new(T) mengalokasikan memori untuk T, mengisinya dengan zero value, dan mengembalikan *T
	pInt := new(int)
	fmt.Println("Zero value via new(int):", *pInt) // 0
	*pInt = 42
	fmt.Println("Setelah diubah:", *pInt)          // 42

	// 2. Menggunakan new() untuk struct
	// Semua field otomatis bernilai zero value ("" dan 0)
	db1 := new(Database)
	fmt.Println("db1 via new():", db1) // &{ 0}

	// 3. Menggunakan &T{} (struct literal) — cara paling idiomatik di Go!
	// Mengalokasikan struct sekaligus menginisialisasi field yang diinginkan
	db2 := &Database{
		Name:    "postgres",
		Timeout: 30,
	}
	fmt.Println("db2 via &T{}:", db2) // &{postgres 30}
}
