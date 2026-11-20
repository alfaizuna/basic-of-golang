package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSummarize(t *testing.T) {
	rows := []Row{
		{Total: 100.0},
		{Total: 200.0},
		{Total: 300.0},
	}
	s := Summarize(rows)
	if s.Count != 3 {
		t.Errorf("expected Count 3, got %d", s.Count)
	}
	if s.SumTotal != 600.0 {
		t.Errorf("expected SumTotal 600.0, got %f", s.SumTotal)
	}
	if s.AvgTotal != 200.0 {
		t.Errorf("expected AvgTotal 200.0, got %f", s.AvgTotal)
	}

	empty := Summarize(nil)
	if empty.Count != 0 || empty.SumTotal != 0 || empty.AvgTotal != 0 {
		t.Errorf("expected zero summary for empty slice, got %+v", empty)
	}
}

func TestGroupByCategory(t *testing.T) {
	rows := []Row{
		{Item: "A", Category: "tech"},
		{Item: "B", Category: "office"},
		{Item: "C", Category: "tech"},
	}
	groups := GroupByCategory(rows)
	if len(groups) != 2 {
		t.Fatalf("expected 2 categories, got %d", len(groups))
	}
	if len(groups["tech"]) != 2 {
		t.Errorf("expected 2 items in 'tech', got %d", len(groups["tech"]))
	}
	if len(groups["office"]) != 1 {
		t.Errorf("expected 1 item in 'office', got %d", len(groups["office"]))
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
	want1 := "Summary: 3 items, Total: $3450.00, Avg: $1150.00"
	want2 := "Categories found: 2 (electronics: 2, furniture: 1)"

	if !strings.Contains(got, want1) {
		t.Errorf("output does not contain %q, got:\n%s", want1, got)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("output does not contain %q, got:\n%s", want2, got)
	}
}
