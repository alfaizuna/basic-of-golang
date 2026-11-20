# Methods on Generic Types

Method yang dideklarasikan pada tipe generik secara otomatis memiliki akses ke parameter tipe dari receiver tersebut.

## Contoh

```go
//go:build ignore

package main

import "fmt"

type Box[T any] struct {
	value T
}

// Method receiver menyertakan parameter tipe [T]
func (b *Box[T]) Set(val T) {
	b.value = val
}

func (b *Box[T]) Get() T {
	return b.value
}

func main() {
	intBox := Box[int]{}
	intBox.Set(42)
	fmt.Println("Isi box:", intBox.Get())
}

```

## Penjelasan

1. **Receiver Syntax**: Saat mendefinisikan method pada tipe generik, nama parameter tipe harus disertakan setelah nama tipe receiver: `func (s *Stack[T]) Push(item T)`.
2. **Tidak Boleh Type Parameter Tambahan**: Method di Go tidak boleh mendeklarasikan parameter tipe baru sendiri di luar yang telah dimiliki receiver. Semua parameter tipe method harus berasal dari receiver tipe tersebut.
3. **Pointer Receiver**: Direkomendasikan jika method memodifikasi internal state dari tipe generik tersebut.

## >_ Exercise

Buat struktur data generik `Stack[T any]` dengan field `items []T`.
Implementasikan 3 method pada `*Stack[T]`:
1. `Push(val T)`: menambahkan `val` ke ujung slice `items`.
2. `Pop() (T, bool)`: jika stack kosong kembalikan zero value dan `false`. Jika ada, ambil elemen terakhir, potong slice, lalu kembalikan elemen tersebut dan `true`.
3. `Len() int`: mengembalikan jumlah elemen saat ini.

### Expected output

```
Pop: 3, Len: 2
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Method pada tipe generik menyertakan parameter tipe receiver, contoh: `(s *Stack[T])`.
- Method tidak dapat mendeklarasikan type parameter baru di luar type parameter receiver.
