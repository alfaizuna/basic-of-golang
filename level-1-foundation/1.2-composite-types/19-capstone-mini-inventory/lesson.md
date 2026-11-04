# Capstone: Mini Inventory

Di proyek akhir Course 1.2 ini, kamu akan menggabungkan semua konsep yang telah dipelajari: **Slice, Map, Struct, dan Method** ke dalam sistem manajemen inventaris barang mini.

## Contoh
example.go
```go
package main

import "fmt"

type Item struct {
	Name string
}

type Inventory struct {
	items map[string]Item
}
```

## Penjelasan
Sistem yang akan dibangun:
1. `Item`: merepresentasikan sebuah produk (ID, Nama, Qty, Harga).
2. `Inventory`: mengelola koleksi item menggunakan `map[string]Item`.
3. Konstruktor `NewInventory()`: menginisialisasi map internal dengan aman.
4. Method `AddItem(item Item)`: menambahkan barang ke map.
5. Method `TotalValue() int`: menghitung total nilai seluruh inventaris (`sum(Qty * Price)`).

## >_ Exercise
Buka [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/19-capstone-mini-inventory/main.go):
1. Definisikan `type Item struct`:
   - `ID` (`string`)
   - `Name` (`string`)
   - `Qty` (`int`)
   - `Price` (`int`)
2. Definisikan `type Inventory struct`:
   - `items` (`map[string]Item`)
3. Buat konstruktor `func NewInventory() *Inventory`:
   - Mengembalikan pointer `&Inventory{items: make(map[string]Item)}`
4. Buat method `func (inv *Inventory) AddItem(it Item)`:
   - Menyimpan item ke map dengan kunci `it.ID`: `inv.items[it.ID] = it`
5. Buat method `func (inv Inventory) TotalValue() int`:
   - Lakukan iterasi `for _, item := range inv.items`
   - Hitung total akumulasi `item.Qty * item.Price`
   - Kembalikan totalnya
6. Di `main()`:
   - Buat inventaris: `inv := NewInventory()`
   - Tambahkan dua barang:
     - `inv.AddItem(Item{ID: "ITM01", Name: "Mechanical Keyboard", Qty: 2, Price: 150000})`
     - `inv.AddItem(Item{ID: "ITM02", Name: "Wireless Mouse", Qty: 3, Price: 50000})`
   - Cetak:
     - `fmt.Println("Jumlah Jenis Item:", len(inv.items))`
     - `fmt.Println("Total Nilai Stok: Rp" + fmt.Sprint(inv.TotalValue()))`

Expected output:
```
Jumlah Jenis Item: 2
Total Nilai Stok: Rp450000
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Selamat! Kamu telah mengintegrasikan struct sebagai model data, map sebagai indexed store, serta method untuk mengenkapsulasi logika kalkulasi inventaris. Inilah pondasi utama arsitektur kode di Go.
