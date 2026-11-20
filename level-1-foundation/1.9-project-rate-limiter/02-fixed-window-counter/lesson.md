# Fixed Window Counter

Algoritma Fixed Window Counter membagi waktu ke dalam jendela tetap (misal per menit) dan menghitung jumlah request dalam jendela tersebut.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"time"
)

type FixedWindowLimiter struct {
	limit       int
	window      time.Duration
	count       int
	windowStart time.Time
	now         func() time.Time
}

func main() {
	fmt.Println("Fixed window limiter initialized")
}

```

## Penjelasan

1. **Konsep Fixed Window**: Jendela waktu ditentukan oleh interval (misalnya 1 detik). Selama waktu saat ini masih berada di dalam jendela (`now.Sub(windowStart) < window`), setiap request menaikkan counter.
2. **Kelebihan & Kekurangan**: Sangat hemat memori (hanya 1 integer dan 1 timestamp), tetapi memiliki kelemahan lonjakan di batas jendela (*boundary burst*).

## >_ Exercise

Lengkapi struct `FixedWindowLimiter`:
- `limit int`
- `window time.Duration`
- `count int`
- `windowStart time.Time`
- `now func() time.Time`

Implementasikan method `(l *FixedWindowLimiter) Allow() bool`:
- Jika `l.now().Sub(l.windowStart) < l.window`:
  - Jika `l.count < l.limit`: naikkan `l.count++` dan kembalikan `true`.
  - Jika `l.count >= l.limit`: kembalikan `false`.
- Jika sudah melewati batas jendela, jangan perbarui dulu (akan dibahas di lesson 03).

### Expected output

```
Allowed 1: true
Allowed 2: true
Allowed 3: false
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Fixed Window Counter melacak jumlah request dalam interval waktu tertentu dengan memory footprint minimal.
