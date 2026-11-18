## Hint 1 (arah umum)
Jalankan program dengan `go run main.go` di terminal dan amati stack trace-nya.

Perhatikan baris kode yang ditunjuk oleh `main.go:...`:
- Di `createProfile`: variabel `var p *Profile` memiliki *zero value* `nil` (tidak menunjuk ke memori manapun). Kamu tidak bisa mengisi field `p.Bio` pada pointer yang `nil`.
- Di `loadConfig`: struct `AppConfig` berhasil dibuat, tetapi field `Server` bertipe `*ServerConfig` dan belum diinisialisasi sehingga bernilai `nil`. Mengakses `cfg.Server.Host` langsung memicu panic.

---

## Hint 2 (sintaks yang relevan)
Untuk membuat pointer yang valid ke sebuah struct baru, gunakan operator address-of `&` dengan struct literal:

```go
p := &Profile{
    Bio: bio,
}
```

Atau untuk field nested di `loadConfig`:
```go
cfg.Server = &ServerConfig{
    Host: host,
    Port: port,
}
```

---

## Hint 3 (hampir jawaban)
```go
func createProfile(bio string) *Profile {
	return &Profile{
		Bio: bio,
	}
}

func loadConfig(appName string, host string, port int) *AppConfig {
	return &AppConfig{
		AppName: appName,
		Server: &ServerConfig{
			Host: host,
			Port: port,
		},
	}
}
```
