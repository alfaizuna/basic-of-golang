package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPipelineExecution(t *testing.T) {
	rows := []Row{
		{Item: "A", Category: " TECH ", Price: 10, Quantity: 2},
		{Item: "B", Category: "OFFICE", Price: 5, Quantity: 1},
	}

	p := NewPipeline(
		NormalizeCategories,
		CalculateTotals,
	)
	if p == nil {
		t.Fatal("expected non-nil pipeline")
	}

	out := p.Execute(rows)
	if len(out) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(out))
	}
	if out[0].Category != "tech" || out[0].Total != 20.0 {
		t.Errorf("stage execution failed on row 0: %+v", out[0])
	}

	// Test AddStage
	p.AddStage(FilterMinTotal(15.0))
	outFiltered := p.Execute(rows)
	if len(outFiltered) != 1 || outFiltered[0].Item != "A" {
		t.Errorf("expected 1 item after AddStage filter, got %+v", outFiltered)
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
	want1 := "Pipeline executed: 3 raw -> 1 processed row"
	want2 := "Final row: TX-101 | Laptop | electronics | $2400.00"
	want3 := "Pipeline Summary: 1 items, Total: $2400.00, Avg: $2400.00"

	if !strings.Contains(got, want1) {
		t.Errorf("output does not contain %q, got:\n%s", want1, got)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("output does not contain %q, got:\n%s", want2, got)
	}
	if !strings.Contains(got, want3) {
		t.Errorf("output does not contain %q, got:\n%s", want3, got)
	}
}
