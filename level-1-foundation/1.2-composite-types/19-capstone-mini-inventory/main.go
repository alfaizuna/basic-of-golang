package main

import "fmt"

type Item struct {
	ID    string
	Name  string
	Qty   int
	Price int
}

type Inventory struct {
	items map[string]Item
}

func NewInventory() *Inventory {
	// TODO: inisialisasi &Inventory{items: make(map[string]Item)}
	return &Inventory{
		items: make(map[string]Item),
	}
}

func (inv *Inventory) AddItem(it Item) {
	// TODO: simpan item ke inv.items[it.ID]
	inv.items[it.ID] = it
}

func (inv Inventory) TotalValue() int {
	// TODO: hitung total nilai seluruh inventaris (sum Qty * Price)
	total := 0
	for _, it := range inv.items {
		total += it.Qty * it.Price
	}
	return total
}

func main() {
	// TODO:
	// 1. Inisialisasi inv := NewInventory()
	inv := NewInventory()
	// 2. Tambahkan Item 1: ID: "ITM01", Name: "Mechanical Keyboard", Qty: 2, Price: 150000
	inv.AddItem(Item{ID: "ITM01", Name: "Mechanical Keyboard", Qty: 2, Price: 150000})
	// 3. Tambahkan Item 2: ID: "ITM02", Name: "Wireless Mouse", Qty: 3, Price: 50000
	inv.AddItem(Item{ID: "ITM02", Name: "Wireless Mouse", Qty: 3, Price: 50000})
	fmt.Println("Jumlah Jenis Item:", len(inv.items))
	fmt.Println("Total Nilai Stok: Rp" + fmt.Sprint(inv.TotalValue()))

	_ = fmt.Println
}
