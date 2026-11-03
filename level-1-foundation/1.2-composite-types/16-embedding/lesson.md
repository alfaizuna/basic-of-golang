# Struct Embedding

Go tidak mendukung konsep *inheritance* (pewarisan kelas) ala OOP tradisional. Sebagai gantinya, Go menggunakan **komposisi lewat struct embedding**.

## Contoh
example.go
```go
package main

import "fmt"

type Person struct {
	Name string
}

type Employee struct {
	Person // Embedded tanpa nama variabel
	ID     string
}

func main() {
	e := Employee{Person: Person{Name: "Alice"}, ID: "EMP01"}
	// Field Name "naik" (promoted) ke level e!
	fmt.Println(e.Name, e.ID)
}
```

## Penjelasan
- Struct embedding dilakukan dengan menuliskan tipe struct di dalam struct lain **tanpa memberi nama field** (contoh: `Person` bukan `P Person`).
- **Field Promotion**: Semua field dari inner struct langsung bisa diakses seolah-olah merupakan milik outer struct (`e.Name` langsung tanpa perlu `e.Person.Name`).
- Meskipun promoted, `e.Person.Name` tetap valid jika terjadi tabrakan nama (*name shadowing*).

## >_ Exercise
Buka [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/16-embedding/main.go):
1. Definisikan tipe struct `Engine`:
   - `Horsepower` (`int`)
2. Definisikan tipe struct `Car`:
   - Lakukan embedding struct `Engine` di dalam `Car` (tanpa nama field).
   - Tambahkan field `Brand` (`string`).
3. Di `main()`:
   - Buat variabel `c` bertipe `Car` dengan `Brand: "Tesla"` dan `Engine: Engine{Horsepower: 450}`.
   - Manfaatkan *field promotion* untuk mengakses `Horsepower`:
     `fmt.Printf("%s berkekuatan %d HP\n", c.Brand, c.Horsepower)`

Expected output:
```
Tesla berkekuatan 450 HP
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Gunakan struct embedding untuk menggabungkan perilaku dan data (*composition over inheritance*). Field dari struct yang di-embed akan "naik" (*promoted*) sehingga bisa langsung diakses dari struct luar.
