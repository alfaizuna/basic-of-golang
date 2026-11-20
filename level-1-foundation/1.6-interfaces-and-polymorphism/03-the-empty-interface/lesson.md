# The Empty Interface (`any`)

Interface kosong tidak mendefinisikan method apapun, sehingga semua tipe di Go secara otomatis memenuhinya.

## Contoh

Di Go 1.18+, `any` adalah alias resmi untuk `interface{}`.

```go
package main

import "fmt"

func describe(val any) {
	fmt.Printf("(%v, %T)\n", val, val)
}

func main() {
	var i any = 42
	describe(i)

	i = "halo gopher"
	describe(i)
}
```

## Penjelasan

1. **Definisi**: Interface kosong didefinisikan sebagai `interface{}` atau `any`. Karena interface ini memerlukan 0 method, semua tipe data di Go (int, string, struct, slice, dll.) memenuhi interface ini.
2. **Kapan Digunakan**: Digunakan ketika fungsi perlu menangani nilai dari tipe yang tidak diketahui sebelumnya, seperti `fmt.Println(a ...any)` atau serializer JSON.
3. **Trade-off**: Menggunakan `any` mengorbankan compile-time type safety. Anda kehilangan proteksi tipe statis Go dan harus berhati-hati saat mengakses nilai aslinya.

## >_ Exercise

Lengkapi fungsi `PrintAll(items []any)` yang mencetak tipe dan nilai dari setiap elemen dalam slice `items` dengan format `[Tipe] Nilai`.
Di fungsi `main()`, buat sebuah slice `[]any` yang menampung minimal 3 nilai berbeda: sebuah `int`, `string`, dan `bool`, lalu panggil `PrintAll`.

### Expected output

```
[int] 100
[string] Go is awesome
[bool] true
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `any` adalah alias untuk `interface{}` sejak Go 1.18.
- Interface kosong dapat menyimpan nilai tipe apapun, tetapi gunakanlah secara bijak karena mengurangi keamanan tipe statis.
