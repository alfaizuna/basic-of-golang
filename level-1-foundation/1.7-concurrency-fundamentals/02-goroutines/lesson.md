# Goroutines

Goroutine adalah thread eksekusi ringan yang dikelola oleh Go runtime hanya dengan menambahkan kata kunci `go` sebelum pemanggilan fungsi.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	fmt.Println("Halo dari goroutine:", name)
}

func main() {
	// Menjalankan sayHello secara konkuren
	go sayHello("Gopher")

	// Goroutine fungsi anonim
	go func(msg string) {
		fmt.Println("Anonim goroutine:", msg)
	}("Belajar Go")

	// Beri jeda sejenak agar goroutine sempat berjalan sebelum main exit
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Main selesai")
}

```

## Penjelasan

1. **Kata Kunci `go`**: Menempatkan `go` di depan pemanggilan fungsi (`go fn()`) akan meluncurkan fungsi tersebut di goroutine baru. Pemanggil akan langsung melanjutkan eksekusi baris berikutnya tanpa menunggu `fn()` selesai (non-blocking).
2. **Kematian `main()`**: Fungsi `main()` berjalan di goroutine utama. Jika fungsi `main()` selesai dieksekusi dan kembali (*return*), seluruh program Go langsung dihentikan seketika, dan semua goroutine anak yang sedang berjalan akan dimatikan paksa tanpa menunggu selesai.
3. **Sinkronisasi**: Menunggu goroutine dengan `time.Sleep` adalah trik sementara untuk pembelajaran, namun di kode produksi kita menggunakan channel atau `sync.WaitGroup`.

## >_ Exercise

Buat fungsi `LaunchWorker(id int, out chan<- string)` yang meluncurkan sebuah goroutine.
Di dalam goroutine tersebut, kirimkan format string `"Worker <id> done"` ke channel `out`.
Di fungsi `main()`, gunakan channel tersebut untuk menerima pesan dari worker dan cetak ke layar.

### Expected output

```
Worker 1 done
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Goroutine dibuat dengan mudah menggunakan kata kunci `go`.
- Jika goroutine utama (`main`) selesai, semua goroutine lain otomatis terhenti.
