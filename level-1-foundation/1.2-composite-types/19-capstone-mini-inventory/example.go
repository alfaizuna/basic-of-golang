//go:build ignore

package main

import "fmt"

type Product struct {
	Name  string
	Qty   int
	Price int
}

type Store struct {
	Catalog map[string]Product
}

func NewStore() *Store {
	return &Store{
		Catalog: make(map[string]Product),
	}
}

func (s *Store) AddProduct(id string, p Product) {
	s.Catalog[id] = p
}

func main() {
	store := NewStore()
	store.AddProduct("P01", Product{Name: "Buku Go", Qty: 5, Price: 100000})
	store.AddProduct("P02", Product{Name: "Buku Java", Qty: 5, Price: 100000})


	fmt.Println("Total produk terdaftar:", len(store.Catalog))
	fmt.Printf("Item: %+v\n", store.Catalog["P01"])
	fmt.Printf("Item: %+v\n", store.Catalog["P02"])

}
