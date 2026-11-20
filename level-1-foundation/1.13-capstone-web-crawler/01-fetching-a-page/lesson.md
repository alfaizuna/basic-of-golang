# Fetching a Page & The Fetcher Interface

Mendefinisikan antarmuka abstraksi `Fetcher` untuk mengambil konten web, serta mengimplementasikan `FakeFetcher` untuk pengujian unit deterministik.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
)

type FetchResult struct {
	Body string
	URLs []string
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func main() {
	fmt.Println("Fetcher interface model")
}

```

## Penjelasan

1. **Desain Abstraksi `Fetcher`**:
   `type Fetcher interface { Fetch(url string) (body string, urls []string, err error) }`
   Dengan menggunakan interface, crawler kita dapat diuji dengan mudah menggunakan `FakeFetcher` in-memory tanpa perlu melakukan koneksi internet sungguhan saat testing.
2. **`FakeFetcher`**: Menggunakan `map[string]*FakeResponse` untuk memetakan URL ke konten tiruan dan daftar URL anak.

## >_ Exercise

1. Definisikan interface:
   ```go
   type Fetcher interface {
       Fetch(url string) (string, []string, error)
   }
   ```
2. Buat struct `FakeResponse`:
   - `Body string`
   - `URLs []string`
3. Buat struct `FakeFetcher`:
   - `Data map[string]*FakeResponse`
4. Implementasikan method `(f *FakeFetcher) Fetch(url string) (string, []string, error)`:
   - Jika `url` ditemukan di `f.Data`, kembalikan `res.Body, res.URLs, nil`.
   - Jika tidak ditemukan, kembalikan `"", nil, fmt.Errorf("not found: %s", url)`.

### Expected output

```
Fetched: Home Page with 2 links
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Mengisolasi pengambilan data di balik interface `Fetcher` membuat seluruh arsitektur crawler mudah diuji tanpa dependensi jaringan eksternal.
