//go:build ignore

package main

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"bilangan genap positif", 4, true},
		{"bilangan ganjil positif", 5, false},
		{"nol adalah genap", 0, true},
		{"bilangan ganjil negatif", -3, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// got := IsEven(tc.input)
			// if got != tc.expected { t.Errorf(...) }
		})
	}
}
