# Capstone: Flag Store 1 - Model

Di tiga lesson terakhir Course 1.3 ini, kamu akan membangun sebuah **In-Memory Feature-Flag Store**.
*Feature flags* digunakan di aplikasi modern untuk mengaktifkan atau menonaktifkan fitur secara dinamis tanpa perlu deploy ulang aplikasi.

Pada tahap pertama ini, kamu akan mendefinisikan model struct dan konstruktor store in-memory.

## Contoh
example.go
```go
package main

type ConfigItem struct {
	Key   string
	Value string
}

type ConfigStore struct {
	items map[string]*ConfigItem
}

func NewConfigStore() *ConfigStore {
	return &ConfigStore{
		items: make(map[string]*ConfigItem),
	}
}
```

## Penjelasan
- `Flag`: merepresentasikan sebuah feature flag.
  - `Name` (`string`): nama/identifier unik flag (misal `"new-checkout"`).
  - `Enabled` (`bool`): status aktif/non-aktif fitur.
  - `Rollout` (`int`): persentase pengguna yang melihat fitur (misal `0-100`).
- `Store`: struct yang menyimpan kumpulan flag.
  - `flags` (`map[string]*Flag`): map dari nama flag ke pointer `*Flag`.
  - Mengapa menyimpan `*Flag` dan bukan `Flag` (value)? Karena dengan menyimpan pointer, kita bisa memanipulasi (*update-in-place*) status flag langsung tanpa perlu menimpa map berulang kali!
- `NewStore() *Store`: fungsi konstruktor yang mengembalikan pointer ke `Store` baru dengan map internal yang sudah diinisialisasi menggunakan `make(map[string]*Flag)`.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/15-flag-store-1-model/main.go).

1. Definisikan struct `Flag`:
   - `Name` bertipe `string`
   - `Enabled` bertipe `bool`
   - `Rollout` bertipe `int`
2. Definisikan struct `Store`:
   - `flags` bertipe `map[string]*Flag`
3. Buat fungsi konstruktor `NewStore() *Store`:
   - Mengembalikan pointer ke `Store` baru dengan field `flags` yang sudah diinisialisasi dengan `make(map[string]*Flag)`.
4. Buat method `(s *Store) Count() int`:
   - Mengembalikan jumlah flag yang tersimpan di dalam map `s.flags`. Jika `s == nil` atau `s.flags == nil`, kembalikan `0`.

Expected output ketika `go run main.go`:
```
Store created. Initial flag count: 0
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Menyimpan pointer di dalam map (`map[string]*T`) adalah pola yang sangat umum di Go saat entitas dalam store perlu dimodifikasi langsung di tempat (*mutate in-place*) secara efisien tanpa duplikasi memori.
