package main

import (
	"errors"
	"fmt"
)

// example.go — hanya untuk referensi materi (read-only)

type Product struct {
	ID    int
	Title string
	Price float64
}

// FindProduct mencari produk berdasarkan ID.
// Pola idiomatik: mengembalikan (*Product, error).
// Jika tidak ditemukan, kembalikan (nil, error).
func FindProduct(id int) (*Product, error) {
	if id != 1 {
		return nil, errors.New("product not found")
	}
	return &Product{ID: 1, Title: "Mechanical Keyboard", Price: 89.99}, nil
}

// ProductSummary menerima pointer produk dengan nil-guard.
func ProductSummary(p *Product) string {
	// Guard clause: amankan jika pointer bernilai nil
	if p == nil {
		return "No product information available"
	}
	return fmt.Sprintf("%s ($%.2f)", p.Title, p.Price)
}
