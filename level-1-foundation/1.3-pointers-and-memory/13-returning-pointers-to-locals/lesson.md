# Returning Pointers to Locals

Di bahasa seperti C atau C++, mengembalikan pointer atau alamat dari variabel lokal fungsi adalah kesalahan fatal (*dangling pointer / undefined behavior*) karena memori lokal di stack akan langsung dihancurkan saat fungsi selesai berjalan.

Namun di **Go**, kamu **sangat aman dan lazim** mengembalikan alamat dari variabel lokal!

## Contoh
example.go
```go
package main

import "fmt"

type Account struct {
	Owner   string
	Balance int
}

// newAccount membuat variabel lokal acc dan mengembalikan alamatnya (&acc).
// Compiler Go mendeteksi variabel ini masih dibutuhkan di luar fungsi,
// sehingga secara otomatis mengalokasikannya di heap (escape analysis).
func newAccount(owner string, balance int) *Account {
	acc := Account{
		Owner:   owner,
		Balance: balance,
	}
	return &acc // ✅ Sangat aman di Go!
}

func main() {
	user := newAccount("Budi", 500000)
	fmt.Printf("Owner: %s, Saldo: Rp%d\n", user.Owner, user.Balance)
}
```

## Penjelasan
- Variabel lokal `acc` dideklarasikan di dalam fungsi `newAccount`.
- Di Go, saat compiler melihat alamat variabel lokal diambil dan dikembalikan (`return &acc`), compiler melakukan proses bernama **Escape Analysis**.
- Karena nilai tersebut harus tetap hidup setelah fungsi selesai, Go memindahkannya (*escape*) ke memori **heap**.
- Pola `return &Struct{...}` atau mendeklarasikan variabel lokal lalu `return &v` adalah pola konstruktor (*factory function*) yang paling standar dan idiomatik di Go.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/13-returning-pointers-to-locals/main.go).

Lengkapi dua fungsi konstruktor lokal berikut:
1. Di [`NewCounter(initial int) *int`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/13-returning-pointers-to-locals/main.go#L16-L19):
   - Deklarasikan variabel lokal integer (misal `val := initial`).
   - Kembalikan alamat dari variabel lokal tersebut (`&val`).
2. Di [`NewProfile(username, email string) *UserProfile`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/13-returning-pointers-to-locals/main.go#L21-L24):
   - Inisialisasi struct lokal `UserProfile` dengan field `Username` dan `Email`.
   - Kembalikan pointer ke struct tersebut (bisa dengan `return &p` atau langsung `return &UserProfile{...}`).

Expected output ketika `go run main.go`:
```
Counter: 10
Profile: budi (budi@example.com)
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Mengembalikan pointer ke variabel lokal sepenuhnya aman di Go. Go runtime dan compiler mengurus siklus hidup memori secara otomatis melalui *escape analysis* dan *garbage collection*.
