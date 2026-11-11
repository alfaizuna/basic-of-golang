package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestAddBonus(t *testing.T) {
	scores := []int{10, 20, 30}
	addBonus(scores, 5)
	want := []int{15, 25, 35}
	for i, v := range want {
		if scores[i] != v {
			t.Errorf("addBonus: scores[%d] ingin %d, dapat %d", i, v, scores[i])
		}
	}
}

func TestResetMap(t *testing.T) {
	stock := map[string]int{"apel": 10, "jeruk": 5}
	resetMap(stock)
	if len(stock) != 0 {
		t.Errorf("resetMap: map seharusnya kosong, masih ada %d key: %v", len(stock), stock)
	}
}

func TestMainOutput(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	scores := []int{10, 20, 30}
	addBonus(scores, 5)
	fmt.Println("scores:", scores)

	stock := map[string]int{"apel": 10, "jeruk": 5}
	resetMap(stock)
	fmt.Println("stock:", stock)

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)

	got := strings.TrimSpace(buf.String())
	if !strings.Contains(got, "scores: [15 25 35]") {
		t.Errorf("output tidak mengandung 'scores: [15 25 35]', dapat:\n%s", got)
	}
	if !strings.Contains(got, "stock: map[]") {
		t.Errorf("output tidak mengandung 'stock: map[]', dapat:\n%s", got)
	}
}
