# Resetting Windows

Menangani transisi waktu: ketika waktu berjalan melewati durasi jendela, window start harus diatur ulang dan counter di-reset kembali ke nol.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Resetting window logic")
}

```

## Penjelasan

1. **Deteksi Kedaluwarsa**: Ketika `currentTime.Sub(l.windowStart) >= l.window`, jendela aktif telah kedaluwarsa.
2. **Reset State**:
   - Setel `l.windowStart = currentTime`.
   - Setel `l.count = 0`.
3. **Penyuntikan Waktu Deterministik**: Dengan mengubah fungsi `now`, kita dapat mensimulasikan lompatan waktu (misal 5 detik kemudian) tanpa menunggu waktu nyata berlalu.

## >_ Exercise

Perbarui method `(l *FixedWindowLimiter) Allow() bool`:
1. Ambil waktu saat ini: `cur := l.now()`.
2. Jika `cur.Sub(l.windowStart) >= l.window`, reset jendela:
   `l.windowStart = cur`
   `l.count = 0`
3. Jika `l.count < l.limit`:
   `l.count++`
   kembalikan `true`.
4. Jika tidak, kembalikan `false`.

### Expected output

```
Req 1: true
Req 2: false
After window reset: true
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Resetting window memastikan kuota request segar kembali setelah durasi jendela berakhir.
