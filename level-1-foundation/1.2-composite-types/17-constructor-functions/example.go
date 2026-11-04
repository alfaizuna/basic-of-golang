//go:build ignore

package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

// Konvensi konstruktor di Go: New<NamaStruct>
func NewBankAccount(owner string, initialBalance float64) (*BankAccount, error) {
	if owner == "" {
		return nil, errors.New("nama pemilik akun tidak boleh kosong")
	}
	if initialBalance < 0 {
		return nil, errors.New("saldo awal tidak boleh negatif")
	}

	return &BankAccount{
		Owner:   owner,
		Balance: initialBalance,
	}, nil
}

func main() {
	acc, err := NewBankAccount("Alice", 100000)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Akun berhasil dibuat: %s (Saldo: Rp%.0f)\n", acc.Owner, acc.Balance)

	// Uji validasi gagal
	_, errInvalid := NewBankAccount("", 500)
	if errInvalid != nil {
		fmt.Println("Validasi gagal:", errInvalid)
	}
}
