# Token Bucket Algorithm

Algoritma Token Bucket menambahkan token ke dalam ember penampung dengan laju konstan, mendukung lonjakan trafik sesaat (*burst*) secara elegan.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	capacity   float64
	tokens     float64
	refillRate float64 // token per detik
	lastRefill time.Time
}

func main() {
	fmt.Println("Token bucket algorithm")
}

```

## Penjelasan

1. **Mekanisme Token Bucket**:
   - Ember memiliki kapasitas maksimum `capacity`.
   - Token diisi ulang (*refill*) berdasarkan selang waktu: `deltaTokens = elapsedSeconds * refillRate`.
   - Jumlah token tidak pernah melebihi `capacity` (`tokens = min(capacity, tokens + delta)`).
   - Setiap request membutuhkan 1 token. Jika `tokens >= 1.0`, kurangi 1 token dan izinkan (`true`). Jika kurang dari 1.0, tolak (`false`).
2. **Kelebihan**: Sangat populer di industri (dipakai oleh AWS, Stripe, Google) karena hemat memori dan mendukung penanganan burst secara mulus.

## >_ Exercise

Buat struct `TokenBucket`:
- `capacity float64`
- `tokens float64`
- `refillRate float64` (token per detik)
- `lastRefill time.Time`
- `now func() time.Time`

Implementasikan method `(tb *TokenBucket) Allow() bool`:
1. Hitung `elapsed := tb.now().Sub(tb.lastRefill).Seconds()`.
2. Update token: `tb.tokens += elapsed * tb.refillRate`.
3. Jika `tb.tokens > tb.capacity`, setel `tb.tokens = tb.capacity`.
4. Perbarui `tb.lastRefill = tb.now()`.
5. Jika `tb.tokens >= 1.0`:
   - `tb.tokens -= 1.0`
   - Kembalikan `true`.
6. Jika tidak, kembalikan `false`.

### Expected output

```
Token 1: true
Token 2: true
Token 3: false
After refill: true
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Token Bucket mengakomodasi lonjakan trafik sesaat (*bursting*) hingga batas kapasitas ember.
- Laju pengisian ulang konstan menjamin stabilitas throughput jangka panjang.
