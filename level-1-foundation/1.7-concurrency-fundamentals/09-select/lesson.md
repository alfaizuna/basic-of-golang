# The Select Statement

Statement `select` memungkinkan goroutine menunggu beberapa operasi channel sekaligus secara non-blocking atau multi-way.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- "dari satu"
	}()

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- "dari dua"
	}()

	// Menunggu mana yang siap lebih dulu
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Menerima:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Menerima:", msg2)
		}
	}
}

```

## Penjelasan

1. **Mirip `switch` tapi untuk Channel**: `select` mengevaluasi semua case operasi channel (kirim atau terima).
2. **Karakteristik Pemilihan**:
   - Jika satu case siap, case tersebut dieksekusi.
   - Jika beberapa case siap bersamaan, Go memilih salah satu secara acak (*pseudo-random uniform selection*), sehingga tidak ada starvation.
   - Jika tidak ada case yang siap, `select` akan memblokir sampai minimal satu channel siap (kecuali ada case `default`).

## >_ Exercise

Buat fungsi `FirstResponder(chA, chB <-chan string) string` yang:
Menggunakan `select` untuk mengambil pesan pertama yang tiba dari `chA` atau `chB` dan mengembalikannya ke pemanggil.

### Expected output

```
First arrived: Fast response
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `select` memungkinkan koordinasi dan penerimaan pesan dari berbagai channel sekaligus.
- Jika beberapa channel siap serempak, Go memilih satu case secara acak.
