# Debugging with Prints & Logs

Memanfaatkan verb formatting canggih `%+v` dan `%#v` pada paket `fmt` serta logger bawaan untuk inspeksi variabel.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"log"
)

type Config struct {
	Host string
	Port int
}

func main() {
	cfg := Config{Host: "localhost", Port: 8080}

	// %v: hanya nilai
	fmt.Printf("%%v:  %v\n", cfg)

	// %+v: menyertakan nama field struct
	fmt.Printf("%%+v: %+v\n", cfg)

	// %#v: representasi sintaks Go lengkap
	fmt.Printf("%%#v: %#v\n", cfg)

	log.Println("Pesan log otomatis menyertakan timestamp")
}

```

## Penjelasan

1. **`%+v` (Plus-V)**: Mencetak struct beserta nama setiap fieldnya (`{Host:localhost Port:8080}`), sangat mempermudah membaca data struct besar.
2. **`%#v` (Hash-V)**: Mencetak representasi sintaksis Go lengkap, termasuk nama tipe paketnya (`main.Config{Host:"localhost", Port:8080}`). Sangat bermanfaat untuk membedakan string kosong `""` vs `nil` pointer.
3. **Paket `log`**: Menyediakan logging standar thread-safe dengan format tanggal dan waktu secara otomatis ke standard error (`stderr`).

## >_ Exercise

Buat fungsi `InspectVariable(val any) (fieldFormat, syntaxFormat string)` yang mengembalikan:
- `fieldFormat`: diformat menggunakan `fmt.Sprintf("%+v", val)`
- `syntaxFormat`: diformat menggunakan `fmt.Sprintf("%#v", val)`

### Expected output

```
Field format: {Host:127.0.0.1 Port:9000}
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `%+v` menampilkan nama field struct untuk inspeksi cepat.
- `%#v` menampilkan representasi sintaks Go lengkap yang tak ambigu.
