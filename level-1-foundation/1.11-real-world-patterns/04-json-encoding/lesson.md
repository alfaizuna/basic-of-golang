# JSON Encoding & Struct Tags

Mengubah data struct Go menjadi representasi JSON menggunakan `json.Marshal` dan mengontrol nama kunci JSON dengan struct tags.

## Contoh

```go
//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags,omitempty"`
}

func main() {
	p := Product{ID: 1, Name: "Keyboard", Price: 75.5}
	bytes, _ := json.Marshal(p)
	fmt.Println(string(bytes))
}

```

## Penjelasan

1. **`json.Marshal(v)`**: Mengubah tipe data Go (biasanya struct atau slice) menjadi format byte JSON (`[]byte`).
2. **Struct Tag `json:"key"`**:
   - `json:"user_id"`: nama field di JSON menjadi `user_id`.
   - `json:"omitempty"`: field akan dihilangkan dari JSON jika nilainya adalah zero value (misal `""`, `0`, atau `nil`).
   - `json:"-"`: field diabaikan sepenuhnya dari serialisasi JSON.
3. **Hanya Exported Fields**: `json.Marshal` **hanya** memproses field struct yang diawali huruf kapital (Exported). Field berhuruf kecil akan diabaikan.

## >_ Exercise

Buat struct `User`:
- `ID int` dengan tag `json:"id"`
- `Username string` dengan tag `json:"username"`
- `Email string` dengan tag `json:"email,omitempty"`

Buat fungsi `ToJSON(u User) (string, error)` yang mengonversi `User` ke string JSON menggunakan `json.Marshal`.

### Expected output

```
{"id":1,"username":"gopher"}
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Struct tags `json:"..."` memetakan nama field Go ke format API JSON.
- Opsi `omitempty` menyembunyikan field yang bernilai zero value.
