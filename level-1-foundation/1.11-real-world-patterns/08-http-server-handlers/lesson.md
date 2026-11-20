# HTTP Server Handlers

Membangun web server dasar di Go menggunakan fungsi handler `http.HandleFunc` dan `http.ResponseWriter`.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Halo Pengunjung!")
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	fmt.Println("Server running on :8080...")
	// http.ListenAndServe(":8080", nil)
}

```

## Penjelasan

1. **Signature Handler**: `func(w http.ResponseWriter, r *http.Request)` adalah signature standar handler web di Go.
2. **`w http.ResponseWriter`**: Objek untuk menulis respon balik ke klien:
   - `w.WriteHeader(code)`: mengirim status code HTTP (misal 200, 404, 500). Wajib dipanggil sebelum menulis body jika statusnya bukan 200.
   - `w.Write(bytes)`: menulis body payload.
3. **`r *http.Request`**: Berisi seluruh data permintaan dari klien (URL, headers, body, method, query parameters).

## >_ Exercise

Buat fungsi handler `PingHandler(w http.ResponseWriter, r *http.Request)` yang:
1. Menyetel status code `http.StatusOK` (200).
2. Menulis teks string `"pong"` ke `w`.

### Expected output

```
pong
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `http.ResponseWriter` dan `*http.Request` adalah dua pilar penanganan web HTTP di Go.
- `httptest.NewRecorder` memungkinkan pengujian unit handler HTTP tanpa perlu membuka port jaringan asli.
