# Channels Basics

Channel adalah saluran komunikasi berpengetikan (*typed conduit*) yang memungkinkan goroutine bertukar data dan melakukan sinkronisasi secara aman tanpa shared memory yang rentan race condition.

## Contoh

```go
//go:build ignore

package main

import "fmt"

func main() {
	// Membuat unbuffered channel bertipe string
	messages := make(chan string)

	// Goroutine pengirim (sender)
	go func() {
		messages <- "ping" // Mengirim data ke channel
	}()

	// Goroutine penerima (receiver di goroutine main)
	msg := <-messages // Menerima data dari channel (blocking hingga pengirim siap)
	fmt.Println("Diterima:", msg)
}

```

## Penjelasan

1. **Motto Go**: *"Do not communicate by sharing memory; instead, share memory by communicating."*
2. **Unbuffered Channel**: Channel yang dibuat tanpa menentukan kapasitas (`make(chan T)`). Karakteristik utamanya adalah **sinkron**:
   - Operasi pengiriman (`ch <- val`) akan memblokir (*block*) sampai ada goroutine lain yang siap menerima (`<-ch`).
   - Operasi penerimaan (`<-ch`) akan memblokir sampai ada goroutine lain yang siap mengirim nilai.
3. **Deadlock**: Jika goroutine mencoba mengirim/menerima dari unbuffered channel tanpa ada pasangan goroutine yang berlawanan, Go runtime akan mendeteksi deadlock dan melakukan panic.

## >_ Exercise

Buat fungsi `ComputeSquare(n int) int` yang:
1. Membuat unbuffered channel integer `ch := make(chan int)`.
2. Meluncurkan goroutine yang menghitung `n * n` dan mengirim hasilnya ke `ch`.
3. Menerima nilai kuadrat dari `ch` dan mengembalikannya.

### Expected output

```
Kuadrat dari 9 adalah 81
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Unbuffered channel bersifat sinkron: operasi send dan receive saling memblokir hingga keduanya siap.
- Channel bertindak sebagai mekanisme sinkronisasi dan transfer data sekaligus.
