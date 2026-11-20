# Generics Best Practices

Prinsip utama Go: sederhanakan sebelum memperluas. Gunakan generics jika menyelesaikan duplikasi kode nyata, jangan gunakan jika interface biasa sudah cukup.

## Contoh

```go
//go:build ignore

package main

import "fmt"

// BAIK: Generics pada struktur data koleksi universal
type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

func (s *Set[T]) Add(val T) {
	s.elements[val] = struct{}{}
}

func main() {
	s := NewSet[string]()
	s.Add("Go")
	fmt.Println("Set berhasil dibuat")
}

```

## Penjelasan

1. **Kapan Menggunakan Generics**:
   - Struktur data umum (Tree, Graph, Queue, Concurrent Map, Cache).
   - Fungsi utilitas slice/map universal (Sort, Filter, Keys, Values).
   - Menghindari kode berulang antar beberapa tipe numerik dasar.
2. **Kapan JANGAN Menggunakan Generics**:
   - Jika interface standar (seperti `io.Reader` atau `http.Handler`) sudah cukup jelas.
   - Hanya untuk mempersingkat 1 fungsi sederhana tanpa repetisi.
   - Jika implementasi method berbeda untuk setiap tipe (gunakan polimorfisme interface reguler).

## >_ Exercise

Buat fungsi pembantu generic slice sederhana: `ToSlice[T any](items ...T) []T` yang mengubah sekumpulan argumen variadik bertipe apapun menjadi sebuah slice baru berpengetikan statis `T`.

### Expected output

```
Slice: [10 20 30]
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Generics sangat bermanfaat untuk tipe koleksi dan pustaka utilitas.
- Utamakan kesederhanaan dan keterbacaan kode; jangan memaksakan generics bila interface biasa sudah memadai.
