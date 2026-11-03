//go:build ignore

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Employee melakukan embed terhadap Person (tanpa nama field)
type Employee struct {
	Person // Embedded struct
	Salary int
	Role   string
}

func main() {
	emp := Employee{
		Person: Person{Name: "Budi", Age: 30},
		Salary: 15000000,
		Role:   "Software Engineer",
	}

	// Field dari Person "naik" (promoted) ke Employee!
	// Kita bisa panggil emp.Name langsung, tidak harus emp.Person.Name
	fmt.Printf("%s (%s) berumur %d tahun\n", emp.Name, emp.Role, emp.Age)
	fmt.Printf("Gaji: Rp%d\n", emp.Salary)
}
