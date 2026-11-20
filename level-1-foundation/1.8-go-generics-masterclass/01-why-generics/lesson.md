# Why Generics

Sebelum Go 1.18, developer terpaksa menduplikasi kode untuk setiap tipe data atau menggunakan `any` (`interface{}`) yang mengorbankan keamanan tipe saat kompilasi.

## Contoh

```go
//go:build ignore

package main

import "fmt"

// Tanpa generics, fungsi harus dibuat terpisah untuk setiap tipe:
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Min int:", MinInt(10, 20))
	fmt.Println("Min float:", MinFloat(3.14, 2.71))
}

```

## Penjelasan

1. **Masalah Duplikasi Kode**: Logika seperti mencari nilai minimum, membalikkan slice, atau membuat struktur data (stack, queue) sama persis algoritmanya, namun tipe datanya berbeda.
2. **Kelemahan `any`**: Menggunakan `interface{}` / `any` menghilangkan type checking saat compile time dan memaksa kita melakukan runtime type assertion yang rentan bug panic dan menimbulkan alokasi memori tambahan (*heap boxing*).
3. **Generics sebagai Solusi**: Generics memungkinkan kita menulis kode satu kali yang bekerja untuk berbagai tipe dengan tetap mempertahankan proteksi kompilator sepenuhnya.

## >_ Exercise

Buat fungsi `MinInt(a, b int) int` dan `MinFloat(a, b float64) float64`.
Keduanya mengembalikan nilai terkecil antara `a` dan `b`.
Panggil kedua fungsi tersebut di `main()` dan cetak hasilnya.

### Expected output

```
Min int: 10
Min float: 2.71
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Menulis fungsi terpisah untuk setiap tipe menimbulkan duplikasi kode (*boilerplate*).
- Menggunakan `any` mengorbankan compile-time type safety; Generics menyelesaikan kedua masalah ini secara elegan.
