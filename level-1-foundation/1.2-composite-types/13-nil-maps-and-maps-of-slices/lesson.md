# Nil Maps & Maps of Slices

Membaca dari nil map tidak menyebabkan error, namun mencoba menulis ke nil map akan menyebabkan **panic**. Selain itu, kombinasi `map[K][]V` adalah struktur data yang sangat sering digunakan untuk pengelompokan data.

## Contoh
example.go
```go
package main

import "fmt"

func main() {
	var m map[string]int // nil map
	fmt.Println(m["key"]) // Aman: mengembalikan 0

	// m["key"] = 10 // PANIC! Menulis ke nil map dilarang.

	// Map of slices
	groups := make(map[string][]string)
	groups["A"] = append(groups["A"], "Alice")
	fmt.Println(groups["A"]) // [Alice]
}
```

## Penjelasan
- Nil map (`var m map[K]V`) aman dibaca dan aman diperiksa panjangnya (`len(m) == 0`).
- **Fatal Trap**: Menulis ke nil map (`m[k] = v`) langsung menghasilkan runtime panic: `panic: assignment to entry in nil map`. Gunakan selalu `make(map[K]V)` sebelum menulis.
- `map[string][]T` dapat langsung dikombinasikan dengan `append`: `groups[k] = append(groups[k], item)`. Jika key belum ada, `groups[k]` otomatis bernilai nil slice yang langsung aman untuk di-append.

## >_ Exercise
Buka [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/13-nil-maps-and-maps-of-slices/main.go):
1. Lengkapi fungsi `addStudent(classroom map[string][]string, grade string, name string)`:
   - Tambahkan `name` ke dalam slice `classroom[grade]` menggunakan `append`.
2. Di `main()`:
   - Inisialisasi map aman `classroom := make(map[string][]string)`.
   - Tambahkan siswa:
     - `addStudent(classroom, "Kelas A", "Budi")`
     - `addStudent(classroom, "Kelas A", "Siti")`
     - `addStudent(classroom, "Kelas B", "Andi")`
   - Cetak:
     - `fmt.Println("Kelas A:", classroom["Kelas A"])`
     - `fmt.Println("Kelas B:", classroom["Kelas B"])`

Expected output:
```
Kelas A: [Budi Siti]
Kelas B: [Andi]
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Jangan pernah menulis ke nil map tanpa inisialisasi `make()`. Saat membuat `map[K][]V`, manfaatkan fakta bahwa membaca key baru menghasilkan nil slice yang bisa langsung diteruskan ke `append`.
