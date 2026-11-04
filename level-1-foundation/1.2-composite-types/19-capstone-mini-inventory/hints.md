### Hint 1
Definisi struct dan konstruktor:
```go
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
	return &Inventory{
		items: make(map[string]Item),
	}
}
```

### Hint 2
Method `AddItem` dan `TotalValue`:
```go
func (inv *Inventory) AddItem(it Item) {
	inv.items[it.ID] = it
}

func (inv Inventory) TotalValue() int {
	total := 0
	for _, it := range inv.items {
		total += it.Qty * it.Price
	}
	return total
}
```

### Hint 3
Kerangka kode lengkap:
```go
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
	return &Inventory{items: make(map[string]Item)}
}

func (inv *Inventory) AddItem(it Item) {
	inv.items[it.ID] = it
}

func (inv Inventory) TotalValue() int {
	total := 0
	for _, it := range inv.items {
		total += it.Qty * it.Price
	}
	return total
}

func main() {
	inv := NewInventory()
	inv.AddItem(Item{ID: "ITM01", Name: "Mechanical Keyboard", Qty: 2, Price: 150000})
	inv.AddItem(Item{ID: "ITM02", Name: "Wireless Mouse", Qty: 3, Price: 50000})

	fmt.Println("Jumlah Jenis Item:", len(inv.items))
	fmt.Println("Total Nilai Stok: Rp" + fmt.Sprint(inv.TotalValue()))
}
```
