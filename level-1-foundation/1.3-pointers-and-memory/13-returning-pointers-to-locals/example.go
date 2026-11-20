package main

import "fmt"

// example.go — hanya untuk referensi, jangan diubah

type Account struct {
	Owner   string
	Balance int
}

func newAccount(owner string, balance int) *Account {
	acc := Account{
		Owner:   owner,
		Balance: balance,
	}
	return &acc
}

var _ = fmt.Sprintf
