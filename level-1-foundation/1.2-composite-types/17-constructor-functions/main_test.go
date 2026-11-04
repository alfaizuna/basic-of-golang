package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestNewProduct(t *testing.T) {
	_, err := NewProduct("", 100)
	if err == nil {
		t.Errorf("NewProduct(\"\", 100) harus mengembalikan error")
	}

	_, err2 := NewProduct("Mouse", 0)
	if err2 == nil {
		t.Errorf("NewProduct(\"Mouse\", 0) harus mengembalikan error")
	}

	p, err3 := NewProduct("Keyboard", 500000)
	if err3 != nil || p == nil || p.Name != "Keyboard" {
		t.Errorf("NewProduct valid gagal: %v, product: %v", err3, p)
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

	expected := "Produk: Laptop, Harga: Rp12000000, Stok: 0\nError: nama produk wajib diisi"
	if output != expected {
		t.Errorf("Output main() tidak sesuai.\nExpected:\n%s\n\nActual:\n%s", expected, output)
	}
}
