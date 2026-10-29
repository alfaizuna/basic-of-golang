### Hint 1
Gunakan `fmt.Sprintf(...)` untuk membuat string terformat dengan kombinasi verb seperti `%q` (string dikutip), `%.2f` (float 2 desimal), dan `%d` (integer).

### Hint 2
Bentuk fungsi `formatProduct` dan `describeValue`:
```go
func formatProduct(name string, price float64, stock int) string {
	return fmt.Sprintf("Product: %q | Price: $%.2f | Stock: %d", name, price, stock)
}

func describeValue(val int) string {
	return fmt.Sprintf("Type: %T | Padded: %04d", val, val)
}
```

### Hint 3
Di dalam fungsi `main()`:
```go
p := formatProduct("Mechanical Keyboard", 89.9, 15)
fmt.Println(p)

d := describeValue(42)
fmt.Println(d)
```
