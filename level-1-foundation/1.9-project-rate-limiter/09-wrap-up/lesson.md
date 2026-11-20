# Wrap-Up: RateLimiter Interface

Menyatukan semua algoritma rate limiter ke dalam interface polimorfik `RateLimiter` tunggal yang dapat disubstitusi sesuai kebutuhan sistem.

## Contoh

```go
//go:build ignore

package main

import "fmt"

type RateLimiter interface {
	Allow(key string) bool
}

type AlwaysAllowLimiter struct{}

func (a *AlwaysAllowLimiter) Allow(key string) bool {
	return true
}

func main() {
	var limiter RateLimiter = &AlwaysAllowLimiter{}
	fmt.Println("Allowed:", limiter.Allow("test"))
}

```

## Penjelasan

1. **Abstraksi Bersih**: Interface `RateLimiter` yang hanya memiliki satu method `Allow(key string) bool` memungkinkan backend berpindah strategi limiter (misal dari Token Bucket ke Fixed Window) tanpa mengubah kode handler HTTP.
2. **Perbandingan Karakteristik**:
   - **Fixed Window**: Paling hemat memori, cocok untuk kuota harian/bulanan.
   - **Sliding Window Log**: Akurasi waktu tertinggi, cocok untuk limit ketat finansial/keamanan dengan volume sedang.
   - **Token Bucket**: Paling fleksibel, standar industri untuk API publik berkinerja tinggi.

## >_ Exercise

1. Definisikan interface:
   ```go
   type RateLimiter interface {
       Allow(key string) bool
   }
   ```
2. Buat struct `SimulatedLimiter` yang mengimplementasikan `RateLimiter`. Struct ini memiliki `maxRequests int` dan `current int`. Jika `current < maxRequests`, naikkan `current` dan kembalikan `true`.
3. Buat fungsi evaluasi `ExecuteRequests(rl RateLimiter, key string, count int) int` yang memanggil `rl.Allow(key)` sebanyak `count` kali dan mengembalikan total request yang berhasil diizinkan.

### Expected output

```
Total permitted: 5 of 10
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Polimorfisme interface menyatukan berbagai strategi pembatasan laju di bawah satu kontrak API yang konsisten.
- Arsitektur modular mempermudah pengujian, benchmarking, dan pergantian strategi algoritma.
