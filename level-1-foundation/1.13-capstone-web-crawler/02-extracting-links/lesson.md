# Extracting Links from HTML

Mem-parsing dokumen HTML menggunakan tokenizer `golang.org/x/net/html` dan mengubah URL relatif menjadi absolut dengan `net/url`.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"net/url"
	"strings"
	"golang.org/x/net/html"
)

func main() {
	baseURL, _ := url.Parse("https://example.com/blog/")
	relURL, _ := url.Parse("../about")
	resolved := baseURL.ResolveReference(relURL)
	fmt.Println("Absolut URL:", resolved.String())
}

```

## Penjelasan

1. **Tokenisasi HTML**: Paket `golang.org/x/net/html` menyediakan `html.NewTokenizer(reader)`. Loop `tokenizer.Next()` membaca token demi token: `html.StartTagToken`, `html.EndTagToken`, `html.TextToken`, dll.
2. **Mencari Atribut `href`**: Jika token adalah tag `<a ...>`, telusuri `token.Attr` untuk menemukan atribut bernama `"href"`.
3. **Resolusi URL Absolut**: Link dalam HTML seringkali relatif (`href="/catalog"` atau `href="item?id=1"`). Fungsi `baseURL.ResolveReference(relURL)` secara cerdas menyelesaikan path relatif ke URL absolut yang valid.

## >_ Exercise

Buat fungsi `ExtractLinks(baseStr, htmlContent string) ([]string, error)` yang:
1. Parse `baseStr` dengan `url.Parse(baseStr)`. Jika error, kembalikan `nil, err`.
2. Gunakan `tokenizer := html.NewTokenizer(strings.NewReader(htmlContent))`.
3. Lakukan loop `for` yang memanggil `tt := tokenizer.Next()`. Jika `tt == html.ErrorToken`, hentikan loop.
4. Jika `tt == html.StartTagToken || tt == html.SelfClosingTagToken`, periksa apakah nama tag adalah `"a"`.
5. Jika ya, ambil nilai atribut `"href"`. Parse href tersebut dengan `url.Parse(val)`.
6. Ubah ke absolut dengan `resolved := parsedBase.ResolveReference(parsedRel)`.
7. Tambahkan `resolved.String()` ke slice hasil dan kembalikan.

### Expected output

```
Extracted link: https://store.com/products
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `golang.org/x/net/html` menyediakan parser standar untuk mengekstrak token dari dokumen HTML.
- `url.ResolveReference` menyelesaikan path relatif ke URL absolut yang sah.
