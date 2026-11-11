## Hint 1 (arah umum)
Kamu perlu menulis *method* pada `BankAccount`, bukan function biasa. Ingat perbedaan: function biasa menerima struct sebagai parameter, sedangkan method ditulis dengan *receiver* di antara kata `func` dan nama method.

---

## Hint 2 (sintaks yang relevan)
Untuk membuat pointer receiver, tuliskan:

```go
func (a *BankAccount) NamaMethod(param float64) {
    // ubah a.Balance di sini
}
```

Gunakan `if` di dalam `Withdraw` untuk memeriksa apakah `amount <= a.Balance` sebelum mengurangi.

---

## Hint 3 (hampir jawaban)
```go
func (a *BankAccount) Deposit(amount float64) {
    a.Balance += ___
}

func (a *BankAccount) Withdraw(amount float64) {
    if amount <= ___ {
        a.Balance -= ___
    }
}
```

Isi bagian `___` dengan variabel yang tepat.
