# Type Constraints

Type constraint membatasi himpunan tipe data yang diizinkan untuk digunakan sebagai type parameter.

## Contoh

```go
//go:build ignore

package main

import "fmt"

// comparable mengizinkan tipe apapun yang mendukung perbandingan == dan !=
func IndexOf[T comparable](s []T, target T) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	langs := []string{"Go", "Python", "Rust"}
	fmt.Println("Index of Rust:", IndexOf(langs, "Rust"))
}

```

## Penjelasan

1. **`any` vs `comparable`**:
   - `any`: tipe apapun boleh, tetapi kita tidak bisa melakukan operator pembanding `==` atau `!=` pada `T`.
   - `comparable`: constraint bawaan Go untuk tipe-tipe yang mendukung operator `==` dan `!=` (seperti boolean, numbers, string, pointer, channel, struct yang field-fieldnya comparable).
2. **Compile-time Guarantee**: Jika pemanggil mencoba menggunakan tipe yang tidak memenuhi constraint (misalnya slice atau map untuk `comparable`), compiler langsung memunculkan compile error.

## >_ Exercise

Buat fungsi generik `Contains[T comparable](items []T, target T) bool` yang:
Memeriksa apakah `target` ada di dalam slice `items`. Kembalikan `true` jika ada, atau `false` jika tidak ditemukan.

### Expected output

```
Contains 20: true
Contains 99: false
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `comparable` adalah constraint bawaan yang mengizinkan pemakaian operator `==` dan `!=`.
- Constraint memberikan jaminan kemampuan (*capabilities*) yang dapat dilakukan terhadap type parameter.
