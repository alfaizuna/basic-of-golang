# Cache with TTL Expiry

DNS Resolver asli selalu memanfaatkan cache agar tidak perlu meminta data ke server otoritatif berulang kali. Data di dalam cache hanya valid selama durasi TTL (Time-To-Live).
Pada pelajaran kelima, kita akan membuat `DNSCache` yang mendukung kedaluwarsa record. Agar mudah dites tanpa perlu memanggil `time.Sleep`, kita menggunakan abstraksi clock: `now func() time.Time`.

## Contoh
example.go
```go
package main

import "time"

type CacheEntry struct {
	Record    *Record
	ExpiresAt time.Time
}

type DNSCache struct {
	entries map[string]*CacheEntry
	now     func() time.Time
}

func NewDNSCache() *DNSCache {
	return &DNSCache{
		entries: make(map[string]*CacheEntry),
		now:     time.Now,
	}
}
```

## Penjelasan
- `ExpiresAt`: waktu absolut saat cache kedaluwarsa, dihitung dengan:
  ```go
  expiresAt := c.now().Add(time.Duration(rec.TTL) * time.Second)
  ```
- `c.now`: fungsi pengambil waktu yang secara default berisi `time.Now`. Pada pengujian, fungsi ini dapat diganti dengan mock waktu instan.
- `Get(name string, recordType RecordType) *Record`:
  - Jika entri ada dan `c.now().Before(entry.ExpiresAt)`, record masih segar (*fresh*) dan dikembalikan.
  - Jika sudah kedaluwarsa (`!c.now().Before(entry.ExpiresAt)`), hapus entri dengan `delete(c.entries, key)` dan kembalikan `nil`.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.4-project-dns-resolver/05-cache-with-ttl-expiry/main.go).

1. Definisikan struct `CacheEntry`:
   - `Record` (`*Record`)
   - `ExpiresAt` (`time.Time`)
2. Definisikan struct `DNSCache`:
   - `entries` (`map[string]*CacheEntry`)
   - `now` (`func() time.Time`)
3. Buat konstruktor `NewDNSCache() *DNSCache`:
   - Inisialisasi `entries: make(map[string]*CacheEntry)`.
   - Inisialisasi `now: time.Now`.
4. Buat method `(c *DNSCache) Set(rec *Record)`:
   - Jika `c == nil` atau `rec == nil`, return.
   - Buat key string unik: `key := string(rec.Type) + ":" + rec.Name`.
   - Hitung `expiresAt := c.now().Add(time.Duration(rec.TTL) * time.Second)`.
   - Simpan `&CacheEntry{Record: rec, ExpiresAt: expiresAt}` ke `c.entries[key]`.
5. Buat method `(c *DNSCache) Get(name string, recordType RecordType) *Record`:
   - Jika `c == nil` atau `c.entries == nil`, return `nil`.
   - Cek `entry, ok := c.entries[string(recordType)+":"+name]`.
   - Jika `ok`:
     - Jika `c.now().Before(entry.ExpiresAt)` (belum expired), return `entry.Record`.
     - Jika sudah expired, hapus dari map dengan `delete(c.entries, key)` dan return `nil`.
   - Return `nil`.

Expected output ketika `go run main.go`:
```
Cached lookup (immediate): 93.184.216.34
Cached lookup (after expiry): nil
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Mengabstraksi pembacaan waktu dengan `now func() time.Time` adalah teknik profesional di Go untuk menguji logika berbasis waktu secara deterministik dan instan.
