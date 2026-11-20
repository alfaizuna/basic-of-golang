# HTTP Client: POST & Timeouts

Mengirim data JSON melalui HTTP POST dan mengonfigurasi batas waktu (timeout) kustom pada `http.Client`.

## Contoh

```go
//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Klien dengan konfigurasi timeout 5 detik
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	payload := []byte(`{"title": "Go Concurrency"}`)
	resp, err := client.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Status:", resp.Status)
}

```

## Penjelasan

1. **Bahaya `http.DefaultClient`**: `http.Get()` dan `http.DefaultClient` **tidak memiliki timeout** bawaan. Jika server tujuan lambat atau macet, koneksi goroutine Anda akan menggantung selamanya.
2. **Membuat Klien dengan Timeout**: Selalu buat instance `&http.Client{Timeout: 10 * time.Second}` di aplikasi backend.
3. **Mengirim Buffer Data**: `bytes.NewBuffer(jsonBytes)` mengubah slice byte menjadi `io.Reader` yang siap dikirim lewat parameter `client.Post`.

## >_ Exercise

Buat fungsi `SendJSONPayload(client *http.Client, url string, jsonData []byte) (int, error)` yang:
1. Mengirim `jsonData` menggunakan `client.Post(url, "application/json", bytes.NewBuffer(jsonData))`.
2. Jika ada error, kembalikan `0, err`.
3. Tutup `resp.Body` via `defer resp.Body.Close()`.
4. Kembalikan status code `resp.StatusCode` dan `nil`.

### Expected output

```
Created status: 201
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Jangan pernah menggunakan `http.Client` tanpa menyetel field `Timeout` di lingkungan produksi.
- `bytes.NewBuffer` menghubungkan byte array ke antarmuka streaming `io.Reader`.
