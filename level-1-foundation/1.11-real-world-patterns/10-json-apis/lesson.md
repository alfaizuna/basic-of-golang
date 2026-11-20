# Building JSON APIs

Membangun REST endpoint yang menerima request JSON dan membalas dengan respon JSON berstandar industri.

## Contoh

```go
//go:build ignore

package main

import (
	"encoding/json"
	"net/http"
)

type CalcRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type CalcResponse struct {
	Result int `json:"result"`
}

func main() {
	// Server JSON handler
}

```

## Penjelasan

1. **Set Header `Content-Type`**: Sebelum memanggil `w.WriteHeader` atau `w.Write`, selalu setel header `w.Header().Set("Content-Type", "application/json")`.
2. **Streaming Decoder/Encoder**:
   - `json.NewDecoder(r.Body).Decode(&req)`: membaca dan mem-parse JSON langsung dari stream request tanpa `io.ReadAll` intermediate.
   - `json.NewEncoder(w).Encode(&res)`: langsung menulis serialisasi JSON ke response stream.
3. **Status Code yang Tepat**: Gunakan `http.StatusBadRequest` (400) untuk format request salah, dan `http.StatusOK` (200) atau `http.StatusCreated` (201) untuk sukses.

## >_ Exercise

Buat struct `SumRequest`:
- `A int` dengan tag `json:"a"`
- `B int` dengan tag `json:"b"`

Buat struct `SumResponse`:
- `Sum int` dengan tag `json:"sum"`

Buat handler `SumHandler(w http.ResponseWriter, r *http.Request)`:
1. Set header `Content-Type: application/json`.
2. Decode request body ke `SumRequest`. Jika ada error, set status `http.StatusBadRequest` (400) dan return.
3. Buat `SumResponse{Sum: req.A + req.B}`.
4. Set status `http.StatusOK` (200) dan encode response ke `w`.

### Expected output

```
{"sum":30}
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `json.NewDecoder` dan `json.NewEncoder` mengalirkan data JSON langsung dari/ke stream HTTP dengan alokasi memori minimal.
- Selalu deklarasikan header `Content-Type: application/json` pada REST API.
