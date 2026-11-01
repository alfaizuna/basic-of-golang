# Presence Checks

Di Go, mengakses key yang tidak ada di map tidak menghasilkan error, melainkan mengembalikan *zero value*. Gunakan *comma-ok idiom* untuk membedakan antara nilai nol asli vs key yang memang tidak ada.

## Contoh
example.go
```go
package main

import "fmt"

func main() {
	scores := map[string]int{"Alice": 100, "Bob": 0}

	val, ok := scores["Charlie"]
	if !ok {
		fmt.Println("Charlie tidak terdaftar")
	}
}
```

## Penjelasan
- `val := m[key]` → jika key tidak ada, `val` akan berisi zero value tipe data tersebut (contoh `0` untuk int, `""` untuk string).
- `val, ok := m[key]` → *comma-ok idiom*. Variabel kedua bertipe boolean (`true` jika key ada, `false` jika tidak ada).
- Pola umum: `if val, ok := m[key]; ok { /* gunakan val */ }`.

## >_ Exercise
Buka [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/10-presence-checks/main.go):
1. Lengkapi fungsi `checkStock(inventory map[string]int, item string)`:
   - Gunakan comma-ok idiom `stock, exists := inventory[item]`.
   - Jika `exists` bernilai `true`, cetak: `fmt.Printf("%s tersedia: %d unit\n", item, stock)`
   - Jika `false`, cetak: `fmt.Printf("%s tidak ditemukan di katalog\n", item)`
2. Di `main()`:
   - Buat map `katalog := map[string]int{"Apel": 5, "Mangga": 0}` (Mangga stoknya 0 tapi terdaftar).
   - Panggil `checkStock(katalog, "Mangga")`
   - Panggil `checkStock(katalog, "Jeruk")`

Expected output:
```
Mangga tersedia: 0 unit
Jeruk tidak ditemukan di katalog
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Jangan mengandalkan perbandingan nilai zero value (seperti `val == 0`) untuk mengecek keberadaan key di map. Selalu gunakan *comma-ok idiom* (`val, ok := m[key]`) untuk kepastian.
