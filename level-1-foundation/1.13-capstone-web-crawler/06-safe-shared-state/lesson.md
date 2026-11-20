# Safe Shared State & Concurrency Limit

Melindungi map visited dengan `sync.Mutex` dan membatasi jumlah goroutine simultan menggunakan channel semaphore.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"sync"
)

type SafeTracker struct {
	mu      sync.Mutex
	visited map[string]bool
}

func (s *SafeTracker) CheckAndMark(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.visited[url] {
		return false // sudah pernah dikunjungi
	}
	s.visited[url] = true
	return true // pertama kali
}

func main() {
	tracker := SafeTracker{visited: make(map[string]bool)}
	fmt.Println("First time:", tracker.CheckAndMark("https://a.com"))
	fmt.Println("Second time:", tracker.CheckAndMark("https://a.com"))
}

```

## Penjelasan

1. **Proteksi Mutex Mutlak**: Ketika puluhan goroutine berjalan serentak, pemeriksaan dan pengisian map visited **harus berada dalam satu critical section** yang dilindungi `sync.Mutex`.
2. **Channel Semaphore**: Meluncurkan 100.000 goroutine sekaligus dapat menghabiskan koneksi dan socket sistem. Kita dapat membatasi beban dengan channel berkapasitas (misal kapasitas 10):
   ```go
   sem := make(chan struct{}, 10)
   sem <- struct{}{}        // Acquire token
   go func() {
       defer func() { <-sem }() // Release token
       // lakukan fetch...
   }()
   ```

## >_ Exercise

Buat struct `SafeTracker`:
- `mu sync.Mutex`
- `visited map[string]bool`

Implementasikan method:
1. `NewSafeTracker() *SafeTracker`: menginisialisasi map.
2. `(s *SafeTracker) Visit(url string) bool`:
   - Kunci mutex.
   - Jika `s.visited[url]` bernilai `true`, return `false`.
   - Setel `s.visited[url] = true` dan return `true`.
3. `(s *SafeTracker) Count() int`: mengembalikan jumlah URL yang tersimpan secara thread-safe.

### Expected output

```
Visited: true
Total unique: 1
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `CheckAndMark` atomik dengan `sync.Mutex` menjamin halaman hanya di-crawl tepat satu kali (*exactly-once*).
- Semaphore channel mencegah lonjakan konkurensi tak terkendali.
