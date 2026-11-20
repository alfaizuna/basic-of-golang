# Buffered Channels

Buffered channel memiliki kapasitas tampung internal, memungkinkan pengiriman data tanpa memblokir hingga antrean buffer terisi penuh.

## Contoh

```go
//go:build ignore

package main

import "fmt"

func main() {
	// Buffered channel dengan kapasitas 2
	ch := make(chan string, 2)

	// Pengiriman tidak memblokir karena kapasitas buffer masih muat
	ch <- "pesan 1"
	ch <- "pesan 2"

	fmt.Println("Panjang buffer saat ini:", len(ch))
	fmt.Println("Kapasitas total channel:", cap(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

```

## Penjelasan

1. **Kapasitas Buffer**: Ditentukan pada argumen kedua fungsi `make(chan T, capacity)`.
2. **Karakteristik Pemblokiran**:
   - Operasi pengiriman (`ch <- val`) **hanya memblokir jika buffer sudah penuh**.
   - Operasi penerimaan (`<-ch`) **hanya memblokir jika buffer kosong**.
3. **Fungsi `len` dan `cap`**:
   - `len(ch)`: jumlah elemen yang saat ini belum diambil di dalam buffer.
   - `cap(ch)`: kapasitas maksimum buffer.

## >_ Exercise

Lengkapi fungsi `QueueJobs(jobs []string) (<-chan string, int)`:
1. Buat buffered channel string dengan kapasitas sesuai `len(jobs)`.
2. Masukkan semua elemen dari `jobs` ke dalam channel secara berurutan (tanpa goroutine, karena kapasitas cukup).
3. Kembalikan channel tersebut beserta kapasitasnya (`cap(ch)`).

### Expected output

```
Buffered jobs count: 3
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Buffered channel hanya memblokir pengirim saat buffer penuh dan penerima saat buffer kosong.
- Cocok digunakan untuk antrean kerja (work queues) atau membatasi lonjakan beban (*burst absorption*).
