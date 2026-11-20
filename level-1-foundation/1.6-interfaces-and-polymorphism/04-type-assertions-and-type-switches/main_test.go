package main

import "testing"

func TestFormatData(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{"string type", "golang", "STR: golang"},
		{"int type", 10, "INT: 20"},
		{"int type zero", 0, "INT: 0"},
		{"bool true", true, "BOOL: YES"},
		{"bool false", false, "BOOL: NO"},
		{"float unknown", 3.14, "UNKNOWN"},
		{"slice unknown", []int{1, 2}, "UNKNOWN"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatData(tc.input)
			if got != tc.expected {
				t.Errorf("FormatData(%v) = %q, want %q", tc.input, got, tc.expected)
			}
		})
	}
}
