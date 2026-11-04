package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestInventory(t *testing.T) {
	inv := NewInventory()
	inv.AddItem(Item{ID: "A", Name: "Item A", Qty: 2, Price: 100})
	inv.AddItem(Item{ID: "B", Name: "Item B", Qty: 5, Price: 50})

	if len(inv.items) != 2 {
		t.Errorf("len(inv.items) = %d, want 2", len(inv.items))
	}

	val := inv.TotalValue()
	expectedVal := (2 * 100) + (5 * 50) // 200 + 250 = 450
	if val != expectedVal {
		t.Errorf("TotalValue() = %d, want %d", val, expectedVal)
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

	expected := "Jumlah Jenis Item: 2\nTotal Nilai Stok: Rp450000"
	if output != expected {
		t.Errorf("Output main() tidak sesuai.\nExpected:\n%s\n\nActual:\n%s", expected, output)
	}
}
