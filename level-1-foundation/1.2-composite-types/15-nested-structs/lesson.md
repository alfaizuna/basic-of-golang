# Nested Structs

Field di dalam struct dapat berupa tipe data apa pun, termasuk struct lain, slice, maupun map.

## Contoh
example.go
```go
package main

import "fmt"

type Address struct {
	City string
}

type User struct {
	Name string
	Addr Address
	Tags []string
}

func main() {
	u := User{
		Name: "Siti",
		Addr: Address{City: "Bandung"},
		Tags: []string{"developer", "mentor"},
	}
	fmt.Printf("%s tinggal di %s\n", u.Name, u.Addr.City)
}
```

## Penjelasan
- **Nested Struct**: Menyusun struct di dalam struct memungkinkan kita membuat model data yang hierarkis dan menyerupai struktur JSON atau entitas database dunia nyata.
- Akses field bertingkat dilakukan secara berurutan dengan operator titik: `u.Addr.City`.
- Slice atau map di dalam struct diakses seperti biasa: `u.Tags[0]`.

## >_ Exercise
Buka [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/15-nested-structs/main.go):
1. Definisikan tipe struct `Location`:
   - `City` (`string`)
   - `Country` (`string`)
2. Definisikan tipe struct `Company`:
   - `Name` (`string`)
   - `Loc` (`Location`)
   - `TechStack` (`[]string`)
3. Di `main()`:
   - Buat variabel `c` bertipe `Company`:
     - `Name: "GopherHub"`
     - `Loc: Location{City: "Surabaya", Country: "Indonesia"}`
     - `TechStack: []string{"Go", "Kubernetes"}`
   - Cetak:
     - `fmt.Printf("Perusahaan: %s di %s\n", c.Name, c.Loc.City)`
     - `fmt.Printf("Tech: %s\n", c.TechStack[0])`

Expected output:
```
Perusahaan: GopherHub di Surabaya
Tech: Go
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Gunakan nested struct untuk memodelkan hubungan data bertingkat. Akses field-nya secara berjenjang menggunakan tanda titik (`parent.child.field`).
