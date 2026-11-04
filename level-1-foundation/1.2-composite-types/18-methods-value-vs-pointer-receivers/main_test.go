package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestWalletMethods(t *testing.T) {
	w := Wallet{Balance: 100}
	if w.BalanceInfo() != "Saldo: Rp100" {
		t.Errorf("BalanceInfo() = %q, want %q", w.BalanceInfo(), "Saldo: Rp100")
	}

	w.Deposit(50)
	if w.Balance != 150 {
		t.Errorf("Setelah Deposit(50), Balance = %d, want 150", w.Balance)
	}
}

func TestMainOutput(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	output := strings.TrimSpace(buf.String())

	expected := "Saldo: Rp50000\nSaldo: Rp75000"
	if output != expected {
		t.Errorf("Output main() tidak sesuai.\nExpected:\n%s\n\nActual:\n%s", expected, output)
	}
}
