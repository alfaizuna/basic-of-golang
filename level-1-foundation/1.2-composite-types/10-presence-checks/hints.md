### Hint 1
Di dalam `checkStock`, gunakan percabangan `if stock, exists := inventory[item]; exists { ... } else { ... }`.

### Hint 2
Bentuk fungsi `checkStock`:
```go
if stock, exists := inventory[item]; exists {
    fmt.Printf("%s tersedia: %d unit\n", item, stock)
} else {
    fmt.Printf("%s tidak ditemukan di katalog\n", item)
}
```

### Hint 3
Kerangka kode lengkap:
```go
func checkStock(inventory map[string]int, item string) {
	if stock, exists := inventory[item]; exists {
		fmt.Printf("%s tersedia: %d unit\n", item, stock)
	} else {
		fmt.Printf("%s tidak ditemukan di katalog\n", item)
	}
}

func main() {
	katalog := map[string]int{"Apel": 5, "Mangga": 0}
	checkStock(katalog, "Mangga")
	checkStock(katalog, "Jeruk")
}
```
