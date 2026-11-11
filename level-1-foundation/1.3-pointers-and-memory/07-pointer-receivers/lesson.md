# Pointer Receivers

Selama ini kita menulis function biasa yang menerima pointer sebagai parameter. Sekarang kita akan belajar cara melakukan hal yang sama, tapi langsung di dalam *method* pada sebuah struct.

## Contoh

```go
// example.go
package main

import "fmt"

type Counter struct {
	Value int
}

// Value receiver — TIDAK mengubah Counter asli
func (c Counter) ShowValue() {
	fmt.Println("Value:", c.Value)
}

// Pointer receiver — mengubah Counter asli
func (c *Counter) Increment() {
	c.Value++
}

func main() {
	c := Counter{Value: 0}
	c.Increment()
	c.Increment()
	c.ShowValue() // Value: 2
}
```

## Penjelasan

- `(c Counter) ShowValue()` → **value receiver**: method menerima salinan struct; perubahan di dalam method tidak terlihat di luar.
- `(c *Counter) Increment()` → **pointer receiver**: method menerima pointer ke struct asli; perubahan *terlihat* di pemanggil.
- `c.Increment()` → Go secara otomatis mengambil alamat `c` menjadi `(&c).Increment()` saat dibutuhkan.
- `c.Value++` di dalam pointer receiver sama artinya dengan `(*c).Value++` — Go auto-dereference.

## >_ Exercise

Kamu punya struct `BankAccount` dengan field `Balance float64`. Tambahkan dua method:

1. `Deposit(amount float64)` — pointer receiver yang menambah `Balance`.
2. `Withdraw(amount float64)` — pointer receiver yang mengurangi `Balance`, tapi **hanya jika** saldo mencukupi (jika tidak cukup, abaikan penarikan).

Setelah itu, di `main()` jalankan:

```go
acc := BankAccount{Balance: 100}
acc.Deposit(50)
acc.Withdraw(30)
acc.Withdraw(200) // tidak cukup saldo, diabaikan
fmt.Printf("Balance: %.2f\n", acc.Balance)
```

Expected output:
```
Balance: 120.00
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Gunakan **pointer receiver** saat method perlu mengubah state struct. Gunakan **value receiver** saat method hanya membaca. Go akan otomatis mengambil alamat variabel jika dibutuhkan, jadi pemanggil tidak perlu menulis `(&acc).Deposit(50)`.
