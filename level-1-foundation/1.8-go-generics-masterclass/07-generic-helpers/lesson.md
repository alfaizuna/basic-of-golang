# Generic Helpers: Map & Filter

Fungsi-fungsi utilitas fungsional seperti Map dan Filter kini dapat ditulis secara universal dan aman dengan Go Generics.

## Contoh

```go
//go:build ignore

package main

import "fmt"

// Map mentransformasi setiap elemen slice []T menjadi []U
func Map[T, U any](items []T, f func(T) U) []U {
	result := make([]U, len(items))
	for i, v := range items {
		result[i] = f(v)
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	squared := Map(nums, func(n int) int {
		return n * n
	})
	fmt.Println("Squared:", squared)
}

```

## Penjelasan

1. **Fleksibilitas Dua Parameter Tipe**: Perhatikan `Map[T, U any]`. Tipe input `T` dan tipe output `U` tidak harus sama. Misalnya kita dapat memetakan slice struct ke slice string id.
2. **Paket Standar `slices` dan `maps`**: Di Go 1.21+, pustaka standar menyertakan paket `slices` (seperti `slices.Index`, `slices.Clone`, `slices.Sort`) dan `maps` yang memanfaatkan generics secara internal.
3. **Alokasi Kapasitas**: Saat membuat slice penampung hasil, gunakan `make([]U, len(items))` jika ukuran pasti sama, atau `make([]T, 0, len(items))` untuk Filter guna menghemat alokasi memori.

## >_ Exercise

Buat fungsi generik `Filter[T any](items []T, predicate func(T) bool) []T` yang:
Menyaring elemen-elemen di mana `predicate(item)` menghasilkan `true` dan mengembalikan slice baru berisi elemen tersebut.

### Expected output

```
Filtered evens: [2 4 6]
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Fungsi utilitas fungsional seperti `Map` dan `Filter` dapat dibuat generik tanpa kehilangan tipe statis.
- Parameter tipe dapat memetakan transformasi dari tipe `T` ke tipe lain `U` secara presisi.
