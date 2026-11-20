# Capstone: Flag Store 2 - Find or Create

Pada tahap kedua ini, kita akan menambahkan kemampuan inti dari *Flag Store*: **mencari atau membuat flag secara dinamis** (*find-or-create*).

Ketika aplikasi meminta sebuah flag:
1. Jika flag sudah pernah terdaftar, store akan mengembalikan pointer ke flag tersebut.
2. Jika flag belum ada, store akan secara otomatis membuatnya dengan nilai default, menyimpannya ke dalam map, dan mengembalikan pointernya.

## Contoh
example.go
```go
package main

type Item struct {
	Key   string
	Count int
}

type Cache struct {
	data map[string]*Item
}

// FindOrCreate mencari item berdasarkan key.
// Jika belum ada, buat baru, simpan ke map, dan kembalikan pointernya.
func (c *Cache) FindOrCreate(key string) *Item {
	if it, exists := c.data[key]; exists {
		return it
	}
	newItem := &Item{Key: key, Count: 0}
	c.data[key] = newItem
	return newItem
}
```

## Penjelasan
- Dengan mengembalikan pointer `*Flag`:
  - Pemanggil mendapatkan rujukan langsung ke objek flag yang tersimpan di dalam memori map store.
  - Setiap pemanggilan `FindOrCreate` dengan nama flag yang sama akan mengembalikan **alamat memori yang persis sama**.
- Pola *Find or Create* sangat umum digunakan pada sistem cache, session manager, metrics collector, dan konfigurasi dinamis.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/16-flag-store-2-find-or-create/main.go).

Lengkapi dua method pada `Store`:
1. [`Get(name string) (*Flag, bool)`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/16-flag-store-2-find-or-create/main.go#L30-L34):
   - Cari flag di dalam `s.flags[name]`.
   - Jika ditemukan, kembalikan `(flag, true)`.
   - Jika tidak ditemukan, kembalikan `(nil, false)`.
2. [`FindOrCreate(name string) *Flag`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/16-flag-store-2-find-or-create/main.go#L36-L45):
   - Periksa apakah flag dengan `name` sudah ada menggunakan `s.flags[name]`.
   - Jika sudah ada, langsung kembalikan pointer flag tersebut.
   - Jika belum ada, buat flag baru bertipe `*Flag` dengan default:
     - `Name: name`
     - `Enabled: false`
     - `Rollout: 0`
   - Simpan pointer flag baru tersebut ke dalam `s.flags[name]` dan kembalikan pointernya.

Expected output ketika `go run main.go`:
```
Flag count: 1
Flag 1: dark-mode (enabled: false, rollout: 0%)
Same memory address: true
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Mengembalikan pointer `*T` dari store memungkinkan pemanggil berbagi objek data yang sama. Dua pemanggilan `FindOrCreate` dengan key yang sama akan menghasilkan alamat memori yang identik.
