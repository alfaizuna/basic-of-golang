# Exported vs Unexported

Enkapsulasi di Go ditentukan sepenuhnya oleh huruf pertama pengenal: huruf kapital berarti publik (exported), huruf kecil berarti privat (unexported).

## Contoh

```go
//go:build ignore

package main

import "fmt"

type BankAccount struct {
	Owner   string  // Exported: bisa diakses dari package lain
	balance float64 // Unexported: hanya bisa diakses di dalam package ini
}

func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.balance += amount
	}
}

func (b *BankAccount) Balance() float64 {
	return b.balance
}

func main() {
	acc := &BankAccount{Owner: "Alice"}
	acc.Deposit(150.0)
	fmt.Printf("%s balance: $%.2f\n", acc.Owner, acc.Balance())
}

```

## Penjelasan

1. **Tanpa Kata Kunci `public` atau `private`**: Go menggunakan konvensi huruf pertama yang sangat sederhana dan konsisten:
   - `Name`, `User`, `Start()`: **Exported** (terlihat oleh package luar).
   - `name`, `user`, `calculate()`: **Unexported** (hanya terlihat di level package).
2. **Desain API Kecil**: Praktik terbaik Go adalah mengekspor sesedikit mungkin (*minimal API surface*). Buat field struct unexported dan sediakan method yang divalidasi dengan baik untuk mencegah manipulasi data ilegal.

## >_ Exercise

Buat struct `Wallet`:
- `owner string` (unexported)
- `balance float64` (unexported)

Sediakan constructor dan method:
1. `NewWallet(owner string, initial float64) *Wallet`
2. `(w *Wallet) Deposit(amount float64)` (hanya tambah jika `amount > 0`)
3. `(w *Wallet) Withdraw(amount float64) bool` (jika `amount > 0` dan `w.balance >= amount`: kurangi balance dan return `true`; selain itu return `false`)
4. `(w *Wallet) Balance() float64` (mengembalikan `w.balance`)
5. `(w *Wallet) Owner() string` (mengembalikan `w.owner`)

### Expected output

```
Alice balance: 120.00
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Huruf kapital = publik (exported); huruf kecil = privat (unexported).
- Sembunyikan field internal struct untuk menjaga integritas data dan menjaga API tetap sederhana.
