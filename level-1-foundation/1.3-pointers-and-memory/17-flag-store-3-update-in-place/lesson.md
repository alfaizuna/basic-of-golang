# Capstone: Flag Store 3 - Update in Place

Di tahap akhir Capstone Course 1.3 ini, kamu akan membuktikan kekuatan sesungguhnya dari pointer: **kemampuan untuk mengubah data langsung di tempat (*update in place*) dan membagikan perubahannya ke seluruh bagian aplikasi tanpa duplikasi alokasi memori**.

## Contoh
example.go
```go
package main

type Counter struct {
	Count int
}

func (c *Counter) Increment() {
	c.Count++
}

func updateCount(c *Counter, amount int) {
	c.Count += amount
}
```

## Penjelasan
- Ketika `store.FindOrCreate("fitur")` mengembalikan pointer `*Flag`:
  - Jika kita memanggil method dengan pointer receiver (misal `flag.Enable()`) atau mengubah field langsung (`flag.Rollout = 75`), perubahan tersebut langsung terjadi di memori struct yang tersimpan di dalam map `store.flags`.
  - Jika komponen lain di aplikasi kemudian memanggil `store.Get("fitur")`, komponen tersebut akan langsung melihat perubahan nilai terbaru!
  - Tidak ada duplikasi data, tidak perlu memanggil fungsi `store.Save(flag)` atau menulis balik ke map.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/17-flag-store-3-update-in-place/main.go).

Lengkapi method pada struct `Flag`:
1. [`(f *Flag) Enable()`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/17-flag-store-3-update-in-place/main.go#L20-L24):
   - Jika `f == nil`, abaikan (*nil guard*).
   - Ubah field `f.Enabled = true`.
2. [`(f *Flag) Disable()`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/17-flag-store-3-update-in-place/main.go#L26-L30):
   - Jika `f == nil`, abaikan (*nil guard*).
   - Ubah field `f.Enabled = false`.
3. [`(f *Flag) SetRollout(percentage int) error`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/17-flag-store-3-update-in-place/main.go#L32-L41):
   - Jika `f == nil`, kembalikan `errors.New("flag is nil")`.
   - Validasi: persentase harus berada di antara `0` dan `100` (`percentage < 0 || percentage > 100`). Jika di luar rentang, kembalikan `errors.New("rollout must be between 0 and 100")`.
   - Ubah field `f.Rollout = percentage`.
   - Kembalikan `nil`.

Expected output ketika `go run main.go`:
```
Initial: beta-ui (enabled: false, rollout: 0%)
After update via pointer:
Retrieved from store: beta-ui (enabled: true, rollout: 75%)
Shared memory confirmed: true
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Selamat! Kamu telah menyelesaikan Capstone Course 1.3: Pointers & Memory!
Dengan menggabungkan konsep *struct, pointer receiver, nil guard*, dan *heap sharing*, kamu berhasil membangun in-memory feature-flag store yang cepat, aman dari crash, dan hemat memori tanpa perlu salin-menyalin struct.
