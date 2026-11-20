# Concurrency Safety

Melindungi state map dan counter dari akses paralel goroutine menggunakan `sync.Mutex` serta memvalidasinya dengan race detector Go.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	mu   sync.Mutex
	data map[string]int
}

func (c *ConcurrentMap) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key]++
}

func main() {
	m := ConcurrentMap{data: make(map[string]int)}
	m.Inc("hit")
	fmt.Println("Hit count:", m.data["hit"])
}

```

## Penjelasan

1. **Map Tidak Aman untuk Konkurensi**: Membaca dan menulis ke Go native `map` dari dua goroutine sekaligus akan langsung menyebabkan runtime fatal: *"fatal error: concurrent map writes"*.
2. **Penyematan Mutex**: Sertakan `mu sync.Mutex` di dalam struct limiter.
3. **Lock & Defer Unlock**: Setiap pemanggilan `Allow()` atau `Cleanup()` wajib mengunci `l.mu.Lock()` dan membukanya via `defer l.mu.Unlock()`.

## >_ Exercise

Buat struct `SafeRateLimiter`:
- `mu sync.Mutex`
- `limit int`
- `counts map[string]int`

Implementasikan method:
- `Allow(key string) bool`:
  - Kunci dengan `s.mu.Lock()` dan `defer s.mu.Unlock()`.
  - Jika `s.counts[key] < s.limit`: naikkan `s.counts[key]++` dan return `true`.
  - Jika sudah mencapai limit, return `false`.
- `GetCount(key string) int`: membaca nilai `s.counts[key]` secara thread-safe.

### Expected output

```
Total allowed for client-1: 10
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `sync.Mutex` mencegah crash fatal *concurrent map read and write*.
- Validasi keamanan thread dengan flag `-race` pada testing.
