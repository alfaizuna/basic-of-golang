# Cleaning Up Idle Clients

Mencegah kebocoran memori (memory leak) pada aplikasi berumur panjang dengan membersihkan data klien yang sudah tidak aktif dari dalam map.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[string]time.Time)
	m["clientA"] = time.Now().Add(-10 * time.Minute)
	// delete key dari map
	delete(m, "clientA")
	fmt.Println("Remaining:", len(m))
}

```

## Penjelasan

1. **Risiko Memory Leak**: Dalam sistem publik, ribuan alamat IP atau token satu kali pakai dapat membanjiri map limiter. Tanpa pembersihan, map akan terus membengkak hingga menghabiskan RAM server (*out-of-memory*).
2. **Field `lastSeen`**: Setiap entri klien mencatat waktu interaksi terakhirnya.
3. **Metode Cleanup**: Menelusuri map dan memanggil fungsi bawaan `delete(m, key)` untuk setiap klien yang `now.Sub(lastSeen) > maxIdle`.

## >_ Exercise

Buat struct `ClientInfo`:
- `lastSeen time.Time`

Buat struct `LimiterWithCleanup`:
- `clients map[string]*ClientInfo`
- `now func() time.Time`

Implementasikan method:
1. `Touch(key string)`: update atau buat entri baru dengan `lastSeen = l.now()`.
2. `Cleanup(maxIdle time.Duration) int`: hapus semua klien yang `l.now().Sub(client.lastSeen) > maxIdle` dari map menggunakan `delete`. Kembalikan jumlah klien yang berhasil dihapus.

### Expected output

```
Removed clients: 1
Remaining clients: 1
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Menghapus entri map dengan `delete(map, key)` melepaskan alokasi memori klien idle.
- Rutin pembersihan berkala (*housekeeping sweep*) adalah praktik wajib pada in-memory stateful services.
