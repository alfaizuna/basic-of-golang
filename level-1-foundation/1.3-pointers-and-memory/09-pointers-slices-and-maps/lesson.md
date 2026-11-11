# Pointers, Slices & Maps

Slice dan map sebenarnya sudah "berbagi" data secara internal — tapi ada jebakan yang sering bikin pemula bingung, terutama saat menggunakan `range`.

## Contoh

```go
// example.go
package main

import "fmt"

func doubleAll(nums []int) {
	for i := range nums {
		nums[i] *= 2 // ✅ ubah elemen asli lewat index
	}
}

func main() {
	numbers := []int{1, 2, 3}
	doubleAll(numbers)
	fmt.Println(numbers) // [2 4 6]
}
```

## Penjelasan

- **Slice** menyimpan pointer ke backing array — mengoper slice ke function sudah "berbagi" datanya; **tidak perlu pointer ke slice** untuk mengubah elemen.
- **Map** juga sudah berbagi — function yang menerima map bisa langsung menambah/menghapus key tanpa perlu `*map`.
- **Jebakan `range`**: `for _, v := range nums` — `v` adalah **salinan** nilainya, bukan elemen aslinya. Gunakan `nums[i]` jika ingin mengubah.
- **Assign ulang slice sendiri** (misal `append`) tetap butuh return value atau pointer ke slice, karena header slice (pointer + len + cap) yang berubah.

## >_ Exercise

Di `main.go` sudah ada dua function dengan bug:

1. `addBonus(scores []int, bonus int)` — mencoba menambahkan `bonus` ke tiap elemen, tapi memakai `v` dari `range` (tidak mengubah slice asli).
2. `resetMap(m map[string]int)` — mencoba mengosongkan map dengan membuat map baru, tapi pemanggil tidak melihat perubahannya.

**Tugasmu:**
- Perbaiki `addBonus` agar elemen slice asli benar-benar berubah.
- Perbaiki `resetMap` agar semua key di map asli terhapus — **tanpa mengubah signature function** (tetap `func resetMap(m map[string]int)`).

Expected output:
```
scores: [15 25 35]
stock: map[]
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Slice dan map sudah berbagi data secara internal — tidak perlu pointer untuk mengubah isinya. Jebakan terbesar adalah `for _, v := range` yang menyalin nilai; gunakan index (`s[i]`) untuk mengubah elemen. Untuk mengosongkan map, hapus tiap key satu per satu dengan `delete`.
