# Recover

Fungsi bawaan `recover()` menangkap kepanikan runtime (panic) di dalam deferred function dan mengembalikan kontrol program.

## Contoh

```go
//go:build ignore

package main

import "fmt"

func safeCall() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Berhasil pulih dari panic:", r)
		}
	}()

	fmt.Println("Melakukan operasi berbahaya...")
	panic("koneksi database meledak")
	fmt.Println("Baris ini tidak akan pernah tercapai")
}

func main() {
	safeCall()
	fmt.Println("Program utama tetap hidup dan berjalan normal!")
}

```

## Penjelasan

1. **Hanya di Dalam `defer`**: Pemanggilan `recover()` hanya efektif jika diletakkan langsung di dalam fungsi yang di-`defer`. Di luar defer, `recover()` selalu mengembalikan `nil` dan tidak berefek.
2. **Mengubah Panic Menjadi Error**: Pola umum di web server (seperti router HTTP) adalah memasang middleware recover agar bug di satu handler tidak mematikan seluruh server backend.
3. **Mengembalikan `any`**: `recover()` mengembalikan nilai yang dilewatkan ke `panic(...)` (biasanya bertipe string atau `error`).

## >_ Exercise

Buat fungsi `SafeRun(fn func()) (err error)` yang:
1. Menjalankan fungsi `fn()`.
2. Jika `fn()` mengalami panic, tangkap via `recover()` di dalam `defer` dan setel named return `err = fmt.Errorf("panic recovered: %v", r)`.
3. Jika `fn()` berjalan sukses tanpa panic, kembalikan `nil`.

### Expected output

```
Caught panic: panic recovered: something went wrong
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `recover()` hanya berfungsi di dalam fungsi `defer`.
- Gunakan recover pada boundary penting (seperti HTTP server worker) untuk mencegah seluruh aplikasi crash.
