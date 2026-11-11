### Hint 1
Periksa pointer `if p != nil` di awal fungsi `applyDiscount` agar tidak terjadi nil-pointer dereference jika fungsi menerima nilai nil.

### Hint 2
- Karena adanya auto-dereference di Go, kamu tidak perlu menulis `(*p).Price`. Cukup tulis `p.Price -= discount`.
- Di dalam `main()`, oper alamat dari `item` menggunakan `&item`: `applyDiscount(&item, 75000)`.

### Hint 3
Kerangka kode lengkapnya:
```go
func applyDiscount(p *Product, discount int) {
	if p != nil {
		p.Price -= discount
	}
}

func main() {
	item := Product{Name: "Keyboard", Price: 500000}
	applyDiscount(&item, 75000)
	fmt.Printf("Produk: %s, Harga: Rp%d\n", item.Name, item.Price)
}
```
