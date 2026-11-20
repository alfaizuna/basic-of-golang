# Data Races

Data race terjadi saat dua goroutine atau lebih mengakses lokasi memori yang sama secara bersamaan dan setidaknya salah satu akses adalah operasi penulisan (write).

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup

	// CONTOH DATA RACE: dua goroutine membaca dan menulis ke counter tanpa proteksi
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // Tidak aman untuk konkurensi!
		}()
	}

	wg.Wait()
	fmt.Println("Counter akhir (mungkin salah):", counter)
}

```

## Penjelasan

1. **Bahaya Data Race**: Data race dapat menyebabkan perilaku tak terduga (*undefined behavior*), data korup yang sulit dilacak, atau crash acak dalam aplikasi produksi.
2. **Race Detector Go**: Go menyediakan alat bawaan canggih bernama Race Detector. Cukup tambahkan flag `-race` saat build atau menjalankan test:
   `go run -race main.go` atau `go test -race ./...`
3. **Mencegah Data Race**:
   - Menghindari shared state (menggunakan channel untuk komunikasi).
   - Menggunakan sinkronisasi mutual exclusion (`sync.Mutex`).
   - Menggunakan operasi atomik dari paket `sync/atomic`.

## >_ Exercise

Perbaiki race condition dengan mengumpulkan hasil dari goroutine melalui channel!
Lengkapi fungsi `AccumulateConcurrent(values []int) int`:
1. Buat channel integer berkapasitas `len(values)`.
2. Luncurkan goroutine untuk setiap elemen yang mengirim nilainya ke channel.
3. Kumpulkan dan jumlahkan semua nilai dari channel di goroutine pemanggil dan kembalikan totalnya.

### Expected output

```
Total sum: 15
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Data race terjadi saat akses memori bersamaan tanpa proteksi sinkronisasi.
- Selalu gunakan `go test -race` untuk mendeteksi potensi konflik akses memori.
