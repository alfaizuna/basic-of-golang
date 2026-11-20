# Results Aggregation & End-to-End Crawl

Menyatukan seluruh modul crawler: mengumpulkan hasil penelusuran, menangani URL gagal, dan melakukan validasi end-to-end.

## Contoh

```go
//go:build ignore

package main

import "fmt"

type CrawlReport struct {
	SuccessURLs []string
	FailedURLs  map[string]string
}

func main() {
	fmt.Println("Crawler capstone wrap-up")
}

```

## Penjelasan

1. **Laporan Hasil (`CrawlReport`)**:
   Mencatat URL yang sukses dikunjungi serta daftar kegagalan (misalnya 404 Not Found atau timeout) tanpa menghentikan proses crawling.
2. **Penyatuan Komponen**:
   - Abstraksi `Fetcher`
   - Parser HTML & resolver absolut URL
   - Filter kedalaman & isolasi domain
   - Konkurensi thread-safe dengan `WaitGroup` dan `Mutex`
3. **Pengujian End-to-End**: Menguji seluruh alur dengan `httptest.NewServer` yang mensimulasikan situs web mini.

## >_ Exercise

Buat struct `CrawlReport`:
- `SuccessURLs []string`
- `FailedURLs map[string]string`
- `mu sync.Mutex`

Implementasikan method thread-safe:
1. `(r *CrawlReport) AddSuccess(url string)`
2. `(r *CrawlReport) AddFailure(url string, errMsg string)`

Buat fungsi `RunCrawler(startURL string, maxDepth int, fetcher Fetcher) *CrawlReport` yang menjalankan penelusuran konkuren secara lengkap.

### Expected output

```
Crawled successfully: 2 pages
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Capstone menggabungkan goroutine, channel, mutex, struct, interface, dan error handling ke dalam arsitektur nyata.
- Crawler konkuren yang dirancang dengan baik beroperasi dengan cepat, stabil, dan thread-safe.
