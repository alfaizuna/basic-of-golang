# Sliding Window Log

Algoritma Sliding Window Log memberikan akurasi absolut dengan menyimpan timestamp dari setiap request dan menghapus timestamp yang kedaluwarsa secara dinamis.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"time"
)

type LogLimiter struct {
	limit      int
	window     time.Duration
	timestamps []time.Time
}

func main() {
	fmt.Println("Sliding window log initialized")
}

```

## Penjelasan

1. **Konsep Sliding Log**: Alih-alih mereset total counter, kita menyimpan riwayat waktu setiap request di slice `[]time.Time`.
2. **Pruning (Pemangkasan)**: Setiap kali request baru masuk pada waktu `cur`, hapus semua timestamp yang `<= cur - window`.
3. **Keputusan**: Jika jumlah timestamp tersisa `< limit`, simpan `cur` ke dalam slice dan izinkan (`true`), jika tidak tolak (`false`).
4. **Keunggulan**: Menghilangkan masalah *boundary burst* pada fixed window secara sempurna.

## >_ Exercise

Buat struct `SlidingWindowLog`:
- `limit int`
- `window time.Duration`
- `timestamps []time.Time`
- `now func() time.Time`

Implementasikan method `(s *SlidingWindowLog) Allow() bool`:
1. Dapatkan `cur := s.now()`.
2. Hitung `cutoff := cur.Add(-s.window)`.
3. Filter `s.timestamps` agar hanya menyisakan timestamp yang `> cutoff`.
4. Jika `len(s.timestamps) < s.limit`:
   - Tambahkan `cur` ke `s.timestamps` dengan `append`.
   - Kembalikan `true`.
5. Jika tidak, kembalikan `false`.

### Expected output

```
Req 1: true
Req 2: true
Req 3: false
After slide: true
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Sliding Window Log menghilangkan celah boundary burst dengan melacak timestamp eksak.
- Biaya komputasinya adalah memori tambahan untuk menyimpan slice timestamp.
