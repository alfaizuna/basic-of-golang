package main

import "fmt"

// TODO:
// 1. Definisikan type Location struct { City string; Country string }
type Location struct {
	City string
	Country string
}
// 2. Definisikan type Company struct { Name string; Loc Location; TechStack []string }
type Company struct {
	Name string
	Loc Location
	TechStack []string
}


func main() {
	// TODO:
	// 1. Buat instance Company:
	//    Name: "GopherHub"
	//    Loc: Location City: "Surabaya", Country: "Indonesia"
	//    TechStack: []string{"Go", "Kubernetes"}

	c := Company{
		Name: "GopherHub",
		Loc: Location{
			City: "Surabaya", Country: "Indonesia",
		},
		TechStack: []string{"Go", "Kubernetes"},
	}

	fmt.Printf("Perusahaan: %s di %s\n", c.Name, c.Loc.City)
	fmt.Printf("Tech: %s\n", c.TechStack[0])

	_ = fmt.Println
}
