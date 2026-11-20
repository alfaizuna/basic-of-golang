# Limiter Model & Clock Injection

Mendesain model dasar Rate Limiter dengan teknik clock injection agar pengujian berbasis waktu dapat berjalan 100% deterministik dan instan.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"time"
)

type Limiter struct {
	limit int
	now   func() time.Time // Clock injection
}

func (l *Limiter) Allow() bool {
	return true
}

func main() {
	limiter := &Limiter{
		limit: 10,
		now:   time.Now,
	}
	fmt.Println("Allowed:", limiter.Allow())
}

```

## Penjelasan

1. **Clock Injection**: Menggantungkan kode pada `time.Now()` langsung membuat unit test lambat (karena harus `time.Sleep`) dan rentan gagal acak (*flaky*). Dengan menyimpan field `now func() time.Time`, test dapat menyuntikkan waktu tiruan yang dapat dimajukan seketika.
2. **Method Receiver**: Struct limiter menggunakan pointer receiver `*Limiter` karena akan mengelola mutasi state internal (seperti counter atau map).

## >_ Exercise

Buat struct `Limiter` dengan field:
- `Limit int`
- `Now func() time.Time`

Buat fungsi constructor `NewLimiter(limit int, now func() time.Time) *Limiter` yang menginisialisasi struct tersebut. Jika parameter `now` adalah `nil`, defaultkan ke `time.Now`.
Tambahkan method `(l *Limiter) Allow() bool` yang sementara mengembalikan `true` jika `l.Limit > 0`.

### Expected output

```
Limiter ready: true
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Clock injection (`now func() time.Time`) membuat pengujian berbasis waktu menjadi deterministik tanpa `time.Sleep`.
- Inisialisasi default parameter nil menjaga kemudahan penggunaan API.
