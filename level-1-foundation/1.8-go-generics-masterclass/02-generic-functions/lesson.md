# Generic Functions

Fungsi generik menggunakan daftar parameter tipe `[T any]` yang dideklarasikan setelah nama fungsi.

## Contoh

```go
//go:build ignore

package main

import "fmt"

// First mengembalikan elemen pertama dari slice bertipe apapun
func First[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[0], true
}

func main() {
	numbers := []int{10, 20, 30}
	names := []string{"Alice", "Bob"}

	// Type inference: compiler otomatis mengetahui tipe T dari argumen
	n, _ := First(numbers)
	name, _ := First(names)

	fmt.Printf("Angka: %d, Nama: %s\n", n, name)
}

```

## Penjelasan

1. **Type Parameter `[T any]`**: Ditempatkan dalam kurung siku `[...]` sebelum daftar parameter fungsi. `T` adalah nama tipe (type parameter), dan `any` adalah constraint yang berarti tipe apa saja diperbolehkan.
2. **Type Inference**: Go compiler cukup pintar untuk menyimpulkan tipe `T` dari argumen yang dilewatkan, sehingga Anda tidak wajib menulis `First[int](numbers)`. Cukup tulis `First(numbers)`.
3. **Zero Value**: Jika slice kosong, variabel `var zero T` akan berisi zero value sesuai tipe konkret `T` (misalnya `0` untuk int, `""` untuk string).

## >_ Exercise

Buat fungsi generik `Last[T any](s []T) (T, bool)` yang:
1. Jika slice `s` kosong, mengembalikan zero value dari `T` dan `false`.
2. Jika slice tidak kosong, mengembalikan elemen terakhir `s[len(s)-1]` dan `true`.

### Expected output

```
Last number: 30, Last fruit: cherry
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Deklarasi fungsi generik menggunakan format `func Name[T any](args) returns`.
- Go mendukung type inference otomatis berdasarkan argumen yang dikirim.
