### Hint 1
Embedding dilakukan tanpa menulis nama variabel di struct `Car`:
```go
type Engine struct {
	Horsepower int
}

type Car struct {
	Engine // anonymous field (embedding)
	Brand  string
}
```

### Hint 2
Saat inisialisasi:
```go
c := Car{
	Engine: Engine{Horsepower: 450},
	Brand:  "Tesla",
}
```

### Hint 3
Kerangka kode lengkap:
```go
type Engine struct {
	Horsepower int
}

type Car struct {
	Engine
	Brand string
}

func main() {
	c := Car{
		Engine: Engine{Horsepower: 450},
		Brand:  "Tesla",
	}

	fmt.Printf("%s berkekuatan %d HP\n", c.Brand, c.Horsepower)
}
```
