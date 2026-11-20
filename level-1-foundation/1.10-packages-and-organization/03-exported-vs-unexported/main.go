package main

import "fmt"

type Wallet struct {
	owner   string
	balance float64
}

func NewWallet(owner string, initial float64) *Wallet {
	// TODO: Inisialisasi dan kembalikan *Wallet (pastikan initial tidak negatif)
	return &Wallet{owner: owner, balance: initial}
}

func (w *Wallet) Deposit(amount float64) {
	// TODO: Tambahkan amount ke balance jika amount > 0
}

func (w *Wallet) Withdraw(amount float64) bool {
	// TODO: Kurangi balance jika saldo cukup dan kembalikan true, selain itu false
	return false
}

func (w *Wallet) Balance() float64 {
	// TODO: Kembalikan balance
	return w.balance
}

func (w *Wallet) Owner() string {
	// TODO: Kembalikan owner
	return w.owner
}

func main() {
	w := NewWallet("Alice", 100.0)
	w.Deposit(50.0)
	w.Withdraw(30.0)
	fmt.Printf("%s balance: %.2f\n", w.Owner(), w.Balance())
}
