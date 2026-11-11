package main

import "fmt"

type BankAccount struct {
	Balance float64
}

// TODO: tambahkan method Deposit(amount float64) di sini
// (gunakan pointer receiver agar Balance asli berubah)

func (acc *BankAccount) Deposit(amount float64) {
	acc.Balance += amount
}

// TODO: tambahkan method Withdraw(amount float64) di sini
// (gunakan pointer receiver; hanya kurangi Balance jika saldo mencukupi)

func (acc *BankAccount) Withdraw(amount float64) {
	if amount <= acc.Balance {
		acc.Balance -= amount
	}
}

func main() {
	acc := BankAccount{Balance: 100}
	acc.Deposit(50)
	acc.Withdraw(30)
	acc.Withdraw(200) // tidak cukup saldo, diabaikan
	fmt.Printf("Balance: %.2f\n", acc.Balance)
}

