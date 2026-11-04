### Hint 1
Gunakan `if` untuk memeriksa parameter `name` dan `price`. Jika salah satu tidak memenuhi syarat, return `nil` dan pesan error via `errors.New(...)`.

### Hint 2
Bentuk fungsi `NewProduct`:
```go
if name == "" {
    return nil, errors.New("nama produk wajib diisi")
}
if price <= 0 {
    return nil, errors.New("harga harus lebih besar dari 0")
}
return &Product{Name: name, Price: price, Stock: 0}, nil
```

### Hint 3
Kerangka kode lengkap:
```go
func NewProduct(name string, price float64) (*Product, error) {
	if name == "" {
		return nil, errors.New("nama produk wajib diisi")
	}
	if price <= 0 {
		return nil, errors.New("harga harus lebih besar dari 0")
	}
	return &Product{Name: name, Price: price, Stock: 0}, nil
}

func main() {
	p, err := NewProduct("Laptop", 12000000)
	if err == nil {
		fmt.Printf("Produk: %s, Harga: Rp%.0f, Stok: %d\n", p.Name, p.Price, p.Stock)
	}

	_, errInvalid := NewProduct("", 5000)
	fmt.Println("Error:", errInvalid)
}
```
