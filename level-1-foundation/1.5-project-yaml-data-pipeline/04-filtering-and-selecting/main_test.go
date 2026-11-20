package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestFilter(t *testing.T) {
	rows := []Row{
		{Category: "books", Total: 15.0},
		{Category: "electronics", Total: 200.0},
		{Category: "books", Total: 30.0},
	}

	books := Filter(rows, func(r Row) bool {
		return r.Category == "books"
	})
	if len(books) != 2 {
		t.Errorf("expected 2 books, got %d", len(books))
	}

	expensive := FilterMinTotal(rows, 100.0)
	if len(expensive) != 1 || expensive[0].Total != 200.0 {
		t.Errorf("expected 1 expensive item with Total 200.0, got %+v", expensive)
	}

	catFiltered := FilterByCategory(rows, "electronics")
	if len(catFiltered) != 1 || catFiltered[0].Category != "electronics" {
		t.Errorf("expected 1 electronics item, got %+v", catFiltered)
	}
}

func TestMainOutput(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)

	got := strings.TrimSpace(buf.String())
	want1 := "Filtered category 'electronics': 2 rows"
	want2 := "Filtered min total $500: 2 rows"

	if !strings.Contains(got, want1) {
		t.Errorf("output does not contain %q, got:\n%s", want1, got)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("output does not contain %q, got:\n%s", want2, got)
	}
}
