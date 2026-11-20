# Concurrent Crawling with Goroutines

Meluncurkan goroutine secara konkuren untuk setiap tautan baru dan mengoordinasikan penyelesaiannya menggunakan `sync.WaitGroup`.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

func crawlStub(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Crawling:", url)
	time.Sleep(10 * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup
	urls := []string{"https://site.com/a", "https://site.com/b"}

	for _, u := range urls {
		wg.Add(1)
		go crawlStub(u, &wg)
	}

	wg.Wait()
	fmt.Println("Selesai")
}

```

## Penjelasan

1. **Kecepatan Konkurensi**: Menelusuri 100 halaman web secara sekuensial membutuhkan 100 detik (jika 1 detik per request). Dengan goroutine konkuren, 100 halaman dapat di-crawl dalam 1-2 detik.
2. **Koordinasi dengan `sync.WaitGroup`**: Setiap URL yang akan di-crawl memanggil `wg.Add(1)` dan diakhiri dengan `defer wg.Done()`. Goroutine utama memanggil `wg.Wait()`.

## >_ Exercise

Buat fungsi `CrawlConcurrentMock(urls []string, processFn func(string))` yang:
1. Menggunakan `sync.WaitGroup`.
2. Meluncurkan satu goroutine untuk setiap elemen di `urls`.
3. Di dalam goroutine, panggil `processFn(u)`.
4. Tunggu semua goroutine selesai sebelum kembali.

### Expected output

```
All workers completed
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `go func()` dan `sync.WaitGroup` mengubah crawling I/O-bound menjadi sangat efisien dan cepat.
