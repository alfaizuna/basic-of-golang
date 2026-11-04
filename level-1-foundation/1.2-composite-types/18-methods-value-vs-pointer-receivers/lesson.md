# Methods: Value vs Pointer Receivers

Method adalah fungsi yang menempel pada sebuah tipe data (receiver). Ada dua jenis receiver: **Value Receiver** dan **Pointer Receiver**.

## Contoh
example.go
```go
package main

import "fmt"

type Counter struct {
	Value int
}

// Value receiver (hanya baca data / read-only)
func (c Counter) Get() int {
	return c.Value
}

// Pointer receiver (memutasi / mengubah data)
func (c *Counter) Add(n int) {
	c.Value += n
}
```

## Penjelasan
- Sintaks method: `func (receiver Tipe) NamaMethod(params) ReturnType`.
- **Value Receiver `(c Counter)`**: Method menerima salinan nilai struct. Cocok jika method hanya membaca data (*read-only*) atau struct berukuran kecil.
- **Pointer Receiver `(c *Counter)`**: Method menerima pointer ke struct asli. **Wajib digunakan** jika method perlu memodifikasi state dari struct, atau jika struct berukuran besar sehingga menyalinnya memakan memori.
- Catatan: Go secara otomatis menangani konversi alamat saat memanggil method (kamu bisa memanggil `c.Add(5)` langsung pada variabel nilai `c`, Go akan mengubahnya menjadi `(&c).Add(5)` di balik layar).

## >_ Exercise
Buka [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/18-methods-value-vs-pointer-receivers/main.go):
1. Diberikan struct:
   `type Wallet struct { Balance int }`
2. Buat value receiver method `(w Wallet) BalanceInfo() string`:
   - Mengembalikan string berformat `"Saldo: Rp" + fmt.Sprint(w.Balance)`
3. Buat pointer receiver method `(w *Wallet) Deposit(amount int)`:
   - Menambahkan `amount` ke `w.Balance`
4. Di `main()`:
   - Buat `w := Wallet{Balance: 50000}`
   - Cetak: `fmt.Println(w.BalanceInfo())`
   - Panggil: `w.Deposit(25000)`
   - Cetak: `fmt.Println(w.BalanceInfo())`

Expected output:
```
Saldo: Rp50000
Saldo: Rp75000
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Gunakan **Value Receiver** jika method hanya membaca data. Gunakan **Pointer Receiver** (`*T`) jika method perlu mengubah data struct asli atau untuk menghindari penyalinan struct besar.
