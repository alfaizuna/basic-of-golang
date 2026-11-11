package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBankAccountMethods(t *testing.T) {
	t.Run("deposit increases balance", func(t *testing.T) {
		acc := BankAccount{Balance: 100}
		acc.Deposit(50)
		if acc.Balance != 150 {
			t.Errorf("setelah Deposit(50) dari 100, ingin Balance=150, dapat %.2f", acc.Balance)
		}
	})

	t.Run("withdraw decreases balance when sufficient", func(t *testing.T) {
		acc := BankAccount{Balance: 150}
		acc.Withdraw(30)
		if acc.Balance != 120 {
			t.Errorf("setelah Withdraw(30) dari 150, ingin Balance=120, dapat %.2f", acc.Balance)
		}
	})

	t.Run("withdraw ignored when insufficient", func(t *testing.T) {
		acc := BankAccount{Balance: 120}
		acc.Withdraw(200)
		if acc.Balance != 120 {
			t.Errorf("Withdraw(200) seharusnya diabaikan, Balance tetap 120, dapat %.2f", acc.Balance)
		}
	})

	t.Run("full scenario output", func(t *testing.T) {
		// Tangkap stdout
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		acc := BankAccount{Balance: 100}
		acc.Deposit(50)
		acc.Withdraw(30)
		acc.Withdraw(200)
		fmt.Printf("Balance: %.2f\n", acc.Balance)

		w.Close()
		os.Stdout = old
		var buf bytes.Buffer
		io.Copy(&buf, r)

		got := strings.TrimSpace(buf.String())
		want := "Balance: 120.00"
		if got != want {
			t.Errorf("output salah:\ningin : %q\ndapat : %q", want, got)
		}
	})
}
