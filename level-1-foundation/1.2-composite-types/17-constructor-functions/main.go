package main

import (
	"errors"
	"fmt"
)

type Product struct {
	Name  string
	Price float64
	Stock int
}

func NewProduct(name string, price float64) (*Product, error) {
	// TODO:
	// 1. Jika name == "", return nil, errors.New("nama produk wajib diisi")
	// 2. Jika price <= 0, return nil, errors.New("harga harus lebih besar dari 0")
	// 3. Return &Product{Name: name, Price: price, Stock: 0}, nil

	if name == "" {
		return nil, errors.New("nama produk wajib diisi")
	}

	if price <= 0 {
		return nil, errors.New("harga harus lebih besar dari 0")
	}

	return &Product{Name: name, Price: price, Stock: 0}, nil
}

func main() {
	// TODO:
	// 1. Buat produk: p, err := NewProduct("Laptop", 12000000)
	p, err := NewProduct("Laptop", 12000000)
	// 2. Jika err == nil, cetak: fmt.Printf("Produk: %s, Harga: Rp%.0f, Stok: %d\n", p.Name, p.Price, p.Stock)
	if err == nil {
		fmt.Printf("Produk: %s, Harga: Rp%.0f, Stok: %d\n", p.Name, p.Price, p.Stock)
	}
	// 3. Buat produk invalid: _, errInvalid := NewProduct("", 5000)
	_, errInvalid := NewProduct("", 5000)
	// 4. Cetak: fmt.Println("Error:", errInvalid)
	fmt.Println("Error:", errInvalid)

	_ = errors.New
	_ = fmt.Println
}
