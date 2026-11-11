# Pointers to Structs

Here's your first look at accessing struct fields through pointers using Go's auto-dereference feature and mutating structs in functions.

## Contoh
example.go
```go
//go:build ignore

package main

import "fmt"

type Player struct {
	Name  string
	Level int
}

// levelUp menerima pointer ke struct Player (*Player)
func levelUp(p *Player) {
	if p == nil {
		return
	}
	// Auto-dereference: Go otomatis mengubah p.Level menjadi (*p).Level
	p.Level++
}

func main() {
	hero := Player{Name: "Arthur", Level: 1}
	fmt.Println("Sebelum level up:", hero) // {Arthur 1}

	// Mengirimkan pointer ke struct hero (&hero)
	levelUp(&hero)

	fmt.Println("Setelah level up:", hero) // {Arthur 2}
}
```

## Penjelasan
- Secara teori, untuk mengakses field struct dari pointer `p *Player`, sintaks dereference eksplisitnya adalah `(*p).Level`.
- Tanda kurung wajib karena operator titik (`.`) memiliki presedensi lebih tinggi daripada operator bintang (`*`).
- Namun, Go memberikan kemudahan sintaks bernama **auto-dereference**: kamu cukup menulis `p.Level`, dan compiler Go secara otomatis menerjemahkannya menjadi `(*p).Level`.
- Hal ini berlaku sama untuk membaca (`val := p.Level`) maupun menulis nilai field baru (`p.Level = 2`).

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/06-pointers-to-structs/main.go):

1. Lengkapi fungsi `applyDiscount(p *Product, discount int)`:
   - Periksa jika `p != nil`.
   - Kurangi field `p.Price` sebesar `discount` (manfaatkan auto-dereference: `p.Price -= discount`).
2. Di dalam `func main()`:
   - Diberikan variabel `item := Product{Name: "Keyboard", Price: 500000}`.
   - Panggil fungsi `applyDiscount` dengan mengoper alamat memori dari `item` (`&item`) dan nilai diskon `75000`.
   - Cetak rincian produk setelah diskon:
     ```go
     fmt.Printf("Produk: %s, Harga: Rp%d\n", item.Name, item.Price)
     ```

Expected output:
```
Produk: Keyboard, Harga: Rp425000
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Go menyediakan fitur **auto-dereference** pada pointer ke struct, sehingga kamu bisa langsung mengakses atau mengubah field struct menggunakan notasi titik (`p.Field`) tanpa perlu menulis `(*p).Field`.
