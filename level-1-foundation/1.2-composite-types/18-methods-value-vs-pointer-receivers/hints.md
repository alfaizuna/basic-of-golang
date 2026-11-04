### Hint 1
Format method value receiver:
```go
func (w Wallet) BalanceInfo() string {
	return "Saldo: Rp" + fmt.Sprint(w.Balance)
}
```

### Hint 2
Format method pointer receiver:
```go
func (w *Wallet) Deposit(amount int) {
	w.Balance += amount
}
```

### Hint 3
Kerangka kode lengkap:
```go
func (w Wallet) BalanceInfo() string {
	return "Saldo: Rp" + fmt.Sprint(w.Balance)
}

func (w *Wallet) Deposit(amount int) {
	w.Balance += amount
}

func main() {
	w := Wallet{Balance: 50000}
	fmt.Println(w.BalanceInfo())

	w.Deposit(25000)
	fmt.Println(w.BalanceInfo())
}
```
