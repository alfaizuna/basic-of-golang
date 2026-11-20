# Closing Channels & Range

Menutup channel memberitahukan penerima bahwa tidak akan ada lagi data baru yang dikirim, memungkinkan pembacaan seluruh elemen dengan loop `for range`.

## Contoh

```go
//go:build ignore

package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // Menutup channel

	// Loop for-range akan berhenti otomatis saat channel ditutup dan buffer habis
	for val := range ch {
		fmt.Println("Menerima:", val)
	}

	// Pemeriksaan status channel dengan comma-ok idiom
	val, ok := <-ch
	fmt.Printf("Nilai: %d, Terbuka: %t\n", val, ok)
}

```

## Penjelasan

1. **Aturan Menutup Channel**:
   - Selalu tutup channel dari sisi **pengirim** (*sender*), jangan pernah dari sisi penerima (*receiver*).
   - Mengirim data ke channel yang sudah ditutup akan menyebabkan **panic**.
   - Menutup channel yang sudah tertutup juga akan menyebabkan **panic**.
2. **Comma-ok Idiom**:
   - `val, ok := <-ch`: jika channel terbuka dan mengirim data, `ok == true`.
   - Jika channel sudah ditutup dan semua data sudah habis dibaca, `ok == false` dan `val` adalah zero value tipe data tersebut.
3. **`for range ch`**: Berjalan terus membaca nilai dari channel sampai channel tersebut ditutup.

## >_ Exercise

Buat fungsi `GenerateFibonacci(n int) <-chan int` yang:
1. Membuat channel integer berkapasitas `n`.
2. Meluncurkan goroutine yang menghitung `n` angka pertama deret Fibonacci (mulai dari 0, 1, 1, 2, 3, 5, dst.), mengirim setiap angka ke channel, dan memanggil `close(ch)` di akhir perhitungan.
3. Mengembalikan channel tersebut sebagai receive-only channel (`<-chan int`).

### Expected output

```
Fibonacci: [0 1 1 2 3]
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `close(ch)` menandai akhir aliran data dan memungkinkan konsumen membaca dengan aman via `for range`.
- Hanya pengirim yang berhak menutup channel; pengiriman ke closed channel berakibat fatal (panic).
