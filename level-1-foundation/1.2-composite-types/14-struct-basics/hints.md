### Hint 1
Definisi struct:
```go
type Book struct {
	Title  string
	Author string
	Pages  int
}
```

### Hint 2
Inisialisasi struct literal:
```go
b := Book{
	Title:  "The Go Programming Language",
	Author: "Alan Donovan",
	Pages:  380,
}
```

### Hint 3
Kerangka kode lengkap:
```go
type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	b := Book{
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
		Pages:  380,
	}

	fmt.Printf("%s karya %s (%d halaman)\n", b.Title, b.Author, b.Pages)
}
```
