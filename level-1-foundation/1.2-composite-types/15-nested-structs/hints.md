### Hint 1
Definisikan struct `Location` terlebih dahulu, lalu jadikan tipe data untuk field `Loc` di dalam `Company`.

### Hint 2
Bentuk tipe data:
```go
type Location struct {
	City    string
	Country string
}

type Company struct {
	Name      string
	Loc       Location
	TechStack []string
}
```

### Hint 3
Kerangka kode lengkap:
```go
type Location struct {
	City    string
	Country string
}

type Company struct {
	Name      string
	Loc       Location
	TechStack []string
}

func main() {
	c := Company{
		Name: "GopherHub",
		Loc: Location{
			City:    "Surabaya",
			Country: "Indonesia",
		},
		TechStack: []string{"Go", "Kubernetes"},
	}

	fmt.Printf("Perusahaan: %s di %s\n", c.Name, c.Loc.City)
	fmt.Printf("Tech: %s\n", c.TechStack[0])
}
```
