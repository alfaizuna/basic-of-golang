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
