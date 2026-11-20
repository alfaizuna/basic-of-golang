# Generic Types

Selain fungsi, kita juga dapat membuat struct, slice, atau map generik dengan menyematkan parameter tipe pada deklarasi tipe.

## Contoh

```go
//go:build ignore

package main

import "fmt"

// Pair menampung dua nilai dengan tipe yang bisa berbeda
type Pair[T, U any] struct {
	First  T
	Second U
}

func main() {
	// Inisialisasi struct generik dengan menyebutkan tipenya secara eksplisit
	p1 := Pair[string, int]{First: "Umur", Second: 25}
	p2 := Pair[int, bool]{First: 200, Second: true}

	fmt.Printf("p1: (%s, %d)\n", p1.First, p1.Second)
	fmt.Printf("p2: (%d, %t)\n", p2.First, p2.Second)
}

```

## Penjelasan

1. **Deklarasi Tipe Generik**: Didefinisikan dengan menambahkan daftar parameter tipe setelah nama struct, misalnya `type Pair[T, U any] struct { ... }`.
2. **Inisialisasi**: Saat menginstansiasi tipe generik secara langsung (tanpa fungsi constructor), tipe parameter **harus disebutkan secara eksplisit**, misalnya `Pair[string, int]{...}`.
3. **Constructor Helper**: Seringkali dibuat fungsi constructor seperti `func NewPair[T, U any](first T, second U) Pair[T, U]` agar pemanggil dapat memanfaatkan type inference.

## >_ Exercise

1. Buat tipe struct generik `Pair[T, U any]` dengan dua field publik: `First T` dan `Second U`.
2. Buat fungsi helper `NewPair[T, U any](first T, second U) Pair[T, U]` yang mengembalikan struct `Pair[T, U]`.
3. Di `main()`, buat pasangan `("HTTP", 200)` menggunakan `NewPair` dan cetak nilainya.

### Expected output

```
Pair: HTTP = 200
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Struct dapat dibuat generik dengan menambahkan parameter tipe pada deklarasi struct.
- Fungsi constructor menyederhanakan inisialisasi tipe generik lewat type inference.
