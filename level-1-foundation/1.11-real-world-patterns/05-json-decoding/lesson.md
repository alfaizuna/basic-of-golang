# JSON Decoding

Membaca dan mem-parsing data JSON menjadi struct Go menggunakan `json.Unmarshal`.

## Contoh

```go
//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Success bool   `json:"success"`
}

func main() {
	raw := `{"status": "ok", "code": 200, "success": true}`
	var resp Response

	err := json.Unmarshal([]byte(raw), &resp) // Wajib passing pointer struct
	if err != nil {
		fmt.Println("Gagal decode:", err)
		return
	}
	fmt.Printf("Parsed: %+v\n", resp)
}

```

## Penjelasan

1. **`json.Unmarshal(data, &dest)`**: Mengurai byte JSON ke dalam variabel tujuan.
2. **Wajib Pointer**: Argumen kedua **harus berupa pointer** (`&struct`), karena `Unmarshal` perlu mengubah isi struct pemanggil secara in-place.
3. **Ketahanan Error**: Jika format JSON rusak (*malformed*), `Unmarshal` mengembalikan error sintaks (`*json.SyntaxError`). Field JSON yang tidak ada di struct Go akan diabaikan secara aman tanpa error.

## >_ Exercise

Buat struct `ServerConfig`:
- `Port int` dengan tag `json:"port"`
- `Host string` dengan tag `json:"host"`
- `Debug bool` dengan tag `json:"debug"`

Buat fungsi `ParseConfig(data []byte) (*ServerConfig, error)` yang mem-parsing `data` ke pointer `ServerConfig`.

### Expected output

```
Port: 8080, Host: localhost, Debug: true
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `json.Unmarshal` selalu membutuhkan pointer ke variabel tujuan (`&dest`).
- Ketidakcocokan tipe atau sintaks rusak akan memicu error yang informatif.
