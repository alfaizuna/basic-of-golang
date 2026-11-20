# Sequential Crawl & Visited Set

Menerapkan penelusuran tautan web sekuensial dengan visited set (`map[string]bool`) untuk mencegah perulangan tak berujung (infinite loop).

## Contoh

```go
//go:build ignore

package main

import "fmt"

func main() {
	visited := make(map[string]bool)
	visited["https://example.com"] = true
	fmt.Println("Sudah dikunjungi:", visited["https://example.com"])
}

```

## Penjelasan

1. **Visited Set**: Halaman web seringkali saling merujuk (A -> B -> A). Tanpa mencatat URL yang sudah dikunjungi, crawler akan terjebak dalam perulangan siklus tanpa akhir (*infinite crawl loop*).
2. **Pola Sequential Crawl**:
   - Jika `visited[url]` bernilai `true`, lewati (*skip*).
   - Tandai `visited[url] = true`.
   - Ambil URL anak via `fetcher.Fetch(url)`.
   - Ulangi penelusuran untuk setiap URL anak.

## >_ Exercise

Buat fungsi `CrawlSequential(url string, fetcher Fetcher, visited map[string]bool) []string` yang:
1. Jika `visited[url]` bernilai `true`, kembalikan `nil`.
2. Tandai `visited[url] = true`.
3. Panggil `_, urls, err := fetcher.Fetch(url)`. Jika ada error, kembalikan `[]string{url}`.
4. Kumpulkan `results := []string{url}`.
5. Untuk setiap `u` dalam `urls`, panggil `CrawlSequential(u, fetcher, visited)` secara rekursif dan gabungkan hasilnya ke `results`.
6. Kembalikan `results`.

### Expected output

```
Total crawled pages: 3
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `map[string]bool` bertindak sebagai visited set untuk memotong siklus penelusuran grafik web.
