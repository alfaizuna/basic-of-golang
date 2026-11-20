package main

import "testing"

func TestWalletEncapsulation(t *testing.T) {
	w := NewWallet("Bob", 50)
	if w.Owner() != "Bob" || w.Balance() != 50 {
		t.Fatal("initial wallet values incorrect")
	}

	w.Deposit(25)
	if w.Balance() != 75 {
		t.Errorf("expected 75 after deposit, got %f", w.Balance())
	}

	if !w.Withdraw(20) || w.Balance() != 55 {
		t.Errorf("withdraw failed or incorrect balance")
	}

	if w.Withdraw(100) {
		t.Errorf("overdraft withdraw should return false")
	}
}
