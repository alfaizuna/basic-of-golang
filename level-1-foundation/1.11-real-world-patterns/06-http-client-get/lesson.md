# HTTP Client: GET

Melakukan permintaan HTTP GET menggunakan pustaka standar `net/http` dan memastikan penutupan `resp.Body`.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("Request gagal:", err)
		return
	}
	defer resp.Body.Close() // Wajib ditutup!

	fmt.Println("Status Code:", resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	_ = body
}

```

## Penjelasan

1. **`http.Get(url)`**: Fungsi helper tingkat tinggi untuk request GET.
2. **Kewajiban Menutup `resp.Body`**: Jika `resp.Body.Close()` tidak dipanggil, koneksi TCP yang mendasarinya tidak dapat digunakan kembali (*connection reuse / keep-alive*), yang pada akhirnya akan menyebabkan sistem kehabisan file descriptor (*socket leak*).
3. **Pola Standar**:
   ```go
   resp, err := http.Get(url)
   if err != nil { return err }
   defer resp.Body.Close()
   ```

## >_ Exercise

Buat fungsi `FetchStatusCodeAndBody(url string) (int, string, error)` yang:
1. Memanggil `http.Get(url)`.
2. Jika ada error jaringan/URL, kembalikan `0, "", err`.
3. Pasang `defer resp.Body.Close()`.
4. Baca body dengan `io.ReadAll(resp.Body)`.
5. Kembalikan status code, body sebagai string, dan `nil`.

### Expected output

```
Status: 200
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Selalu periksa `err != nil` sebelum melakukan `defer resp.Body.Close()`.
- Gunakan `httptest.NewServer` untuk menguji panggilan HTTP Client secara lokal dan deterministik.
