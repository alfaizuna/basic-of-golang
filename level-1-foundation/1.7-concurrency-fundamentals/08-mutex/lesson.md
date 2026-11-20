# Mutex & Safe State

sync.Mutex menjamin mutual exclusion sehingga hanya satu goroutine yang dapat mengakses critical section kode pada satu waktu.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"sync"
)

type SafeStore struct {
	mu   sync.Mutex
	data map[string]int
}

func (s *SafeStore) Set(key string, val int) {
	s.mu.Lock()
	defer s.mu.Unlock() // Dilepaskan saat fungsi Set selesai
	s.data[key] = val
}

func (s *SafeStore) Get(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data[key]
}

func main() {
	store := SafeStore{data: make(map[string]int)}
	store.Set("gopher", 100)
	fmt.Println("Stored:", store.Get("gopher"))
}

```

## Penjelasan

1. **Critical Section**: Bagian kode yang mengakses shared resources (seperti map, slice, counter) yang dapat mengalami modifikasi secara bersamaan.
2. **`Lock()` dan `Unlock()`**:
   - `s.mu.Lock()`: memperoleh hak akses eksklusif. Jika goroutine lain sedang memegang lock, eksekusi akan tertahan (*block*).
   - `s.mu.Unlock()`: melepaskan hak akses eksklusif sehingga goroutine lain yang antre dapat melanjutkan.
3. **Idiom `defer mu.Unlock()`**: Menempatkan `defer mu.Unlock()` tepat setelah pemanggilan `Lock()` memastikan lock selalu dilepaskan, bahkan jika fungsi keluar lebih awal (*early return*) atau terjadi error.

## >_ Exercise

Buat struct `SafeCounter` yang memiliki:
- `mu sync.Mutex`
- `value int`

Implementasikan dua method dengan pointer receiver:
1. `(c *SafeCounter) Inc()`: menambah `value` sebesar 1 dengan proteksi mutex.
2. `(c *SafeCounter) Value() int`: membaca dan mengembalikan `value` dengan proteksi mutex.

Di `main()`, luncurkan 50 goroutine yang masing-masing memanggil `Inc()`, gunakan `sync.WaitGroup` untuk menunggu semuanya, lalu cetak nilai akhirnya.

### Expected output

```
Final count: 50
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `sync.Mutex` mengamankan critical section dari eksekusi bersamaan.
- Gunakan `defer mu.Unlock()` tepat setelah `mu.Lock()` untuk mencegah accidental deadlock.
