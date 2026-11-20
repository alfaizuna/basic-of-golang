# Depth Limit & Same-Domain Filtering

Membatasi jelajahan crawler agar hanya menelusuri domain yang sama dan tidak melebihi kedalaman (*max depth*) tertentu.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"net/url"
)

func isSameDomain(targetStr, allowedHost string) bool {
	u, err := url.Parse(targetStr)
	if err != nil {
		return false
	}
	return u.Host == allowedHost
}

func main() {
	fmt.Println("Same domain check:", isSameDomain("https://shop.com/deals", "shop.com"))
}

```

## Penjelasan

1. **Depth Limiting**: Menjelajahi internet tanpa batas kedalaman dapat menghabiskan waktu berhari-hari. Batas kedalaman (misal `depth <= 2`) membatasi pohon rekursi pencarian.
2. **Same-Domain Policy**: Tanpa filter domain, crawler toko Anda akan keluar menelusuri link eksternal (Twitter, Facebook, Google) dan tersesat di domain luar.
3. **Penyaringan URL**: Ekstrak `targetURL.Host` dan pastikan nilainya sama persis dengan `baseHost`.

## >_ Exercise

Buat fungsi `IsAllowed(targetURL, baseHost string, currentDepth, maxDepth int) bool` yang:
1. Kembalikan `false` jika `currentDepth > maxDepth`.
2. Parse `targetURL` dengan `url.Parse`. Jika error, kembalikan `false`.
3. Kembalikan `true` jika `parsed.Host == baseHost`, selain itu `false`.

### Expected output

```
Allowed: true
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `maxDepth` mengontrol radius jelajah crawler.
- Memvalidasi `url.Host` mencegah crawler meninggalkan lingkup domain target.
