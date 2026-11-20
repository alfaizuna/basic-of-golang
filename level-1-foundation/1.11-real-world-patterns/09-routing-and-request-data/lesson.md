# Routing & Request Data

Menggunakan `http.NewServeMux` untuk isolasi routing, validasi HTTP method, dan membaca query parameter.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		fmt.Fprintf(w, "Hasil pencarian untuk: %s", query)
	})
}

```

## Penjelasan

1. **`http.NewServeMux()`**: Hindari penggunaan `http.DefaultServeMux` global di aplikasi produksi untuk mencegah registrasi route liar dari dependensi luar.
2. **Pemeriksaan Method**: Periksa `if r.Method != http.MethodGet` untuk menolak request dengan method yang tidak sesuai.
3. **Membaca Query Param**: `r.URL.Query().Get("name")` mengambil nilai parameter dari URL query string (misal `/greet?name=Budi`).

## >_ Exercise

Buat fungsi handler `UserSearchHandler(w http.ResponseWriter, r *http.Request)`:
1. Periksa `r.Method`. Jika bukan `http.MethodGet`, set status `http.StatusMethodNotAllowed` (405) dan tulis `"Method Not Allowed"`, lalu kembali.
2. Ambil parameter `"q"` dari query string. Jika kosong, setel default ke `"all"`.
3. Set status `http.StatusOK` dan cetak `"Query: <q>"`. (contoh: `"Query: all"`).

### Expected output

```
Query: all
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `http.NewServeMux()` mengisolasi routing server secara rapi.
- Validasi HTTP method dan sanitasi query parameter adalah standar wajib keamanan API.
