# Timeouts & Default in Select

Kombinasi `time.After` dan klausa `default` dalam `select` memungkinkan penanganan timeout dan operasi channel non-blocking.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	// Non-blocking receive dengan default
	select {
	case msg := <-ch:
		fmt.Println("Menerima:", msg)
	default:
		fmt.Println("Channel kosong, lanjutkan tanpa memblokir!")
	}

	// Operasi dengan timeout
	select {
	case msg := <-ch:
		fmt.Println("Menerima:", msg)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("Timeout! Operasi dibatalkan setelah 50ms")
	}
}

```

## Penjelasan

1. **Klausa `default`**: Jika disematkan pada `select`, operasi menjadi **non-blocking**. Jika tidak ada case channel yang siap, blok `default` langsung dieksekusi seketika.
2. **`time.After(duration)`**: Mengembalikan channel `<-chan time.Time` yang akan mengirim sinyal timestamp setelah durasi berlalu.
3. **Mencegah Deadlock & Slow Network**: Sangat penting untuk membatasi durasi panggilan jaringan atau antrean goroutine agar tidak menggantung tanpa batas (*hung indefinitely*).

## >_ Exercise

Buat fungsi `FetchWithTimeout(ch <-chan string, timeout time.Duration) (string, error)` yang:
1. Mendengarkan `ch`. Jika pesan diterima sebelum timeout, kembalikan pesan tersebut dan `nil`.
2. Jika waktu melebihi `timeout` (gunakan `time.After(timeout)`), kembalikan string kosong `""` dan error dengan pesan `"operation timed out"` (gunakan `fmt.Errorf("operation timed out")`).

### Expected output

```
Result: Success!
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `time.After` dalam `select` adalah pola standar Go untuk mencegah blocking tanpa batas.
- `default` membuat `select` langsung melompat tanpa menunggu, berguna untuk non-blocking channel polling.
