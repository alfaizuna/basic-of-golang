package main

import "fmt"

type Wallet struct {
	Balance int
}

func (w Wallet) BalanceInfo() string {
	// TODO: kembalikan "Saldo: Rp" + fmt.Sprint(w.Balance)
	return "Saldo: Rp" + fmt.Sprint(w.Balance)
}

func (w *Wallet) Deposit(amount int) {
	// TODO: tambahkan amount ke w.Balance
	w.Balance += amount
}

func main() {
	// TODO:
	// 1. Buat w := Wallet{Balance: 50000}
	w := Wallet{Balance: 50000}
	// 2. Cetak: fmt.Println(w.BalanceInfo())
	fmt.Println(w.BalanceInfo())
	// 3. Panggil: w.Deposit(25000)
	w.Deposit(25000)
	// 4. Cetak: fmt.Println(w.BalanceInfo())
	fmt.Println(w.BalanceInfo())	
	_ = fmt.Println
}
