# Union Constraints & ~

Operator union `|` menggabungkan beberapa tipe menjadi satu constraint, sedangkan tilde `~` mengikutsertakan tipe-tipe kustom yang memiliki underlying type yang sama.

## Contoh

```go
//go:build ignore

package main

import "fmt"

// Constraint yang menerima tipe bertipe dasar int atau float64
type Number interface {
	~int | ~float64
}

type MyInt int // Tipe kustom dengan underlying type int

func Double[T Number](n T) T {
	return n * 2
}

func main() {
	var a int = 10
	var b MyInt = 20
	var c float64 = 4.5

	fmt.Println(Double(a))
	fmt.Println(Double(b)) // Valid berkat simbol ~
	fmt.Println(Double(c))
}

```

## Penjelasan

1. **Union `|`**: Menentukan daftar tipe yang diizinkan, misalnya `int | int64 | float64`.
2. **Simbol Tilde `~`**: Tanda `~int` berarti *"tipe apapun yang underlying type-nya adalah `int`"*. Tanpa tanda `~`, tipe kustom `type MyID int` akan ditolak oleh compiler karena dianggap berbeda dengan `int`.
3. **Operator yang Diizinkan**: Di dalam fungsi generik, Anda hanya boleh menggunakan operator yang didukung oleh **semua** tipe di dalam union (misalnya operator aritmatika `+`, `-`, `*` didukung oleh semua numeric types).

## >_ Exercise

1. Definisikan type constraint kustom `OrderedNumber`:
   ```go
   type OrderedNumber interface {
       ~int | ~float64
   }
   ```
2. Buat fungsi generik `Max[T OrderedNumber](a, b T) T` yang mengembalikan nilai terbesar antara `a` dan `b`.
3. Buat custom type `type Speed float64` dan uji `Max` di fungsi `main()`.

### Expected output

```
Max speed: 85.50
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Union `|` menggabungkan beberapa tipe menjadi satu constraint yang fleksibel.
- Tilde `~T` sangat penting agar tipe turunan (*defined types*) tetap diterima oleh constraint.
