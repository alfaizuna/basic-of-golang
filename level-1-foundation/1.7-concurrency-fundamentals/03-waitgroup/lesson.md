# WaitGroup

sync.WaitGroup menyediakan cara idiomatis untuk menunggu sekelompok goroutine menyelesaikan tugasnya secara deterministik.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Pastikan Done() dipanggil saat fungsi selesai
	fmt.Printf("Worker %d mulai\n", id)
	time.Sleep(20 * time.Millisecond)
	fmt.Printf("Worker %d selesai\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Tambah counter sebelum meluncurkan goroutine
		go worker(i, &wg) // Wajib passing pointer *sync.WaitGroup
	}

	wg.Wait() // Blokir sampai counter kembali ke 0
	fmt.Println("Semua worker telah selesai")
}

```

## Penjelasan

1. **Tiga Method Utama**:
   - `Add(delta int)`: menambah hitungan goroutine yang ditunggu. Panggil selalu **sebelum** meluncurkan `go func()` agar terhindar dari race condition.
   - `Done()`: mengurangi hitungan sebanyak 1 (ekuivalen dengan `Add(-1)`). Sangat dianjurkan dipanggil via `defer wg.Done()`.
   - `Wait()`: memblokir goroutine pemanggil sampai hitungan di dalam WaitGroup kembali ke 0.
2. **Wajib Pointer**: `sync.WaitGroup` berisi state internal yang tidak boleh disalin (*must not be copied*). Jika dilewatkan ke fungsi lain, selalu gunakan pointer `*sync.WaitGroup`.

## >_ Exercise

Lengkapi fungsi `ProcessBatch(n int) int` yang:
1. Memiliki variabel counter lokal `var total int` dan `var mu sync.Mutex`.
2. Menggunakan `sync.WaitGroup` untuk menjalankan `n` buah goroutine worker.
3. Setiap worker menambahkan angka `10` ke `total` secara aman (gunakan mutex).
4. Menunggu semua worker selesai dengan `wg.Wait()`, lalu mengembalikan nilai `total`.

### Expected output

```
Total batch: 50
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Gunakan `sync.WaitGroup` untuk menunggu sekumpulan goroutine selesai.
- Selalu panggil `wg.Add()` sebelum goroutine dimulai dan `defer wg.Done()` di dalam goroutine.
