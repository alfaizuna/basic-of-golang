# Per-Client Limits

Menerapkan pembatasan laju independen untuk setiap pengguna atau API key menggunakan map of client state.

## Contoh

```go
//go:build ignore

package main

import "fmt"

type ClientState struct {
	count int
}

func main() {
	clients := make(map[string]*ClientState)
	clients["user-123"] = &ClientState{count: 1}
	fmt.Println("Client count:", clients["user-123"].count)
}

```

## Penjelasan

1. **Pemetaan Klien**: Menggunakan `map[string]*ClientState` di mana key adalah identifier (seperti IP address, User ID, atau API token).
2. **Inisialisasi Klien Baru**: Jika key belum ada di map, buat `ClientState` baru untuk klien tersebut.
3. **Isolasi Kuota**: Aktivitas atau limit yang tercapai oleh `userA` tidak akan memengaruhi jatah request `userB`.

## >_ Exercise

Buat struct `ClientState`:
- `count int`
- `windowStart time.Time`

Buat struct `MultiClientLimiter`:
- `limit int`
- `window time.Duration`
- `clients map[string]*ClientState`
- `now func() time.Time`

Implementasikan method `(l *MultiClientLimiter) Allow(clientKey string) bool`:
- Dapatkan atau inisialisasi state klien untuk `clientKey`.
- Lakukan reset window jika sudah kedaluwarsa.
- Evaluasi limit dan naikkan count jika masih diizinkan.

### Expected output

```
User A 1: true
User A 2: false
User B 1: true
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `map[string]*ClientState` memisahkan state counter per identitas klien secara bersih.
- Lazy initialization membuat klien baru dibuat saat request pertama tiba.
