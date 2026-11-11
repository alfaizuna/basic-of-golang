package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

// applyDiscount mengurangi Price dari produk sesuai nominal discount
func applyDiscount(p *Product, discount int) {
	// TODO: cek jika p != nil, lalu kurangi p.Price sebesar discount
	if p != nil {
		p.Price -= discount
	}
}

func main() {
	item := Product{Name: "Keyboard", Price: 500000}

	// TODO: Panggil applyDiscount dengan alamat memori item (&item) dan diskon 75000
	applyDiscount(&item, 75000)

	// TODO: Cetak dengan fmt.Printf("Produk: %s, Harga: Rp%d\n", item.Name, item.Price)
	fmt.Printf("Produk: %s, Harga: Rp%d\n", item.Name, item.Price)
}
