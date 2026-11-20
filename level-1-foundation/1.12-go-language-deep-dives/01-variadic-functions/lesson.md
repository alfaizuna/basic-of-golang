# Variadic Functions

Fungsi variadik dapat menerima nol atau lebih argumen dari tipe yang ditentukan menggunakan elipsis `...T`.

## Contoh

```go
//go:build ignore

package main

import "fmt"

func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println("Sum 3 angka:", Sum(1, 2, 3))

	// Unpacking slice yang sudah ada dengan '...'
	slice := []int{10, 20, 30}
	fmt.Println("Sum slice:", Sum(slice...))
}

```

## Penjelasan

1. **Sintaks `...T`**: Parameter terakhir pada fungsi dapat berupa elipsis diikuti tipe data (`nums ...int`). Di dalam badan fungsi, parameter tersebut diperlakukan sebagai slice `[]T`.
2. **Penyebaran Slice (`s...`)**: Jika Anda sudah memiliki variabel slice, gunakan sintaks suffix `s...` saat memanggil fungsi variadik untuk membedah (*unpack*) slice menjadi argumen individu.
3. **Contoh di Standard Library**: `fmt.Println(a ...any)`, `append(slice, elements...)`.

## >_ Exercise

Buat fungsi variadik `JoinStrings(sep string, words ...string) string` yang menggabungkan seluruh `words` menggunakan pemisah `sep`.
Jika `words` kosong, kembalikan string kosong `""`.
Uji di `main()` dengan memanggil `JoinStrings("-", "go", "is", "fast")`.

### Expected output

```
go-is-fast
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `...T` mengubah argumen berulang menjadi slice di dalam fungsi.
- Teruskan slice ke fungsi variadik dengan menyematkan suffix `...`.
