# Named Return Values

Mendeklarasikan nama pada nilai kembalian fungsi bertindak sebagai variabel lokal otomatis dan mendokumentasikan tujuan return value.

## Contoh

```go
//go:build ignore

package main

import "fmt"

// Named return values: min dan max otomatis terdefinisi sebagai variabel lokal int
func MinMax(a, b int) (min, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return min, max // Eksplisit lebih disukai daripada naked return
}

func main() {
	low, high := MinMax(15, 8)
	fmt.Printf("Min: %d, Max: %d\n", low, high)
}

```

## Penjelasan

1. **Variabel Lokal Zero Value**: Nilai kembalian bernama (`func F() (res int, err error)`) diinisialisasi otomatis dengan zero value di awal eksekusi fungsi.
2. **Dokumentasi Kode**: Sangat membantu keterbacaan ketika fungsi mengembalikan beberapa nilai bertipe sama (misal `(x, y int)` atau `(width, height float64)`).
3. **Naked Return vs Eksplisit**: Statement `return` kosong tanpa argumen (*naked return*) secara otomatis mengembalikan nilai variabel bernama tersebut. Namun, di fungsi berukuran sedang/panjang, naked return sangat dianjurkan dihindari (*code smell*) karena membingungkan pembaca kode.

## >_ Exercise

Buat fungsi `CalculateStats(numbers []int) (min int, max int, total int)`:
- Jika slice kosong, kembalikan `0, 0, 0`.
- Inisialisasi `min` dan `max` dengan `numbers[0]`.
- Hitung nilai minimum, maksimum, dan jumlah total elemen `numbers`.
- Kembalikan secara eksplisit `return min, max, total`.

### Expected output

```
Min: 2, Max: 9, Total: 16
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Named return values mendokumentasikan arti dari setiap nilai kembalian.
- Hindari naked return pada fungsi panjang untuk menjaga kejelasan alur logika.
