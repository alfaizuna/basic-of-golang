# Optional Values

Mari pelajari bagaimana pointer digunakan untuk merepresentasikan nilai opsional (*unset*) dan membedakannya dari *zero value*.

## Contoh
example.go
```go
type NotificationConfig struct {
	EmailAlerts *bool // nil = belum diset (default: true)
	MaxRetries  *int  // nil = belum diset (default: 3)
}

func BoolPtr(b bool) *bool {
	return &b
}

func ResolveNotification(cfg NotificationConfig) ResolvedNotification {
	res := ResolvedNotification{
		EmailAlerts: true, // nilai default
		MaxRetries:  3,    // nilai default
	}
	if cfg.EmailAlerts != nil {
		res.EmailAlerts = *cfg.EmailAlerts
	}
	if cfg.MaxRetries != nil {
		res.MaxRetries = *cfg.MaxRetries
	}
	return res
}
```

## Penjelasan
- Tipe primitif di Go selalu memiliki *zero value* (`0`, `""`, `false`), sehingga sulit membedakan apakah pengguna sengaja memasukkan angka `0` atau field tersebut memang belum diisi.
- Menggunakan pointer (`*int`, `*string`, `*bool`) memungkinkan kita mendeteksi kondisi *unset* melalui nilai `nil`.
- Jika pointer bernilai `nil`, kita terapkan nilai default; jika tidak `nil`, kita dereference (`*ptr`) untuk mengambil nilai yang diberikan.
- Karena Go tidak mengizinkan mengambil alamat literal secara langsung seperti `&"localhost"`, fungsi helper seperti `StringPtr(s string) *string` umum digunakan untuk membuat pointernya.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/12-optional-values/main.go).

Lengkapi dua fungsi berikut:
1. Di [`StringPtr(s string) *string`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/12-optional-values/main.go#L29-L32): kembalikan alamat memori dari parameter `s`.
2. Di [`ResolveServerConfig(cfg ServerConfig) ResolvedConfig`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/12-optional-values/main.go#L39-L52):
   - Periksa masing-masing pointer `cfg.Host`, `cfg.Port`, dan `cfg.Debug`.
   - Jika field tersebut **tidak `nil`**, timpa nilai default di struct `res` dengan nilai dereferensi dari pointer tersebut.

Expected output ketika `go run main.go`:
```
Default Config: localhost:8080 (debug: false)
Custom Config: 0.0.0.0:3000 (debug: true)
Explicit Zero Config: localhost:0 (debug: false)
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Pointer sangat berguna untuk merepresentasikan nilai opsional atau *unset* (`nil`), membedakannya dari *zero value* bawaan tipe data, dan menjadi dasar pola konfigurasi ber-default di Go.
