package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestApplyDiscount(t *testing.T) {
	prod := Product{Name: "Mouse", Price: 200000}
	applyDiscount(&prod, 50000)

	if prod.Price != 150000 {
		t.Errorf("applyDiscount gagal memutasi field Price: dapat %d, diharapkan 150000", prod.Price)
	}
}

func TestApplyDiscountNilPointer(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("applyDiscount(nil, 1000) menyebabkan panic: %v", r)
		}
	}()

	applyDiscount(nil, 1000)
}

func TestMainOutput(t *testing.T) {
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Gagal membuat pipe: %v", err)
	}
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		outC <- buf.String()
	}()

	main()

	_ = w.Close()
	os.Stdout = oldStdout
	out := <-outC

	expected := "Produk: Keyboard, Harga: Rp425000"
	actual := strings.TrimSpace(out)

	if actual != expected {
		t.Errorf("Output belum sesuai.\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}
