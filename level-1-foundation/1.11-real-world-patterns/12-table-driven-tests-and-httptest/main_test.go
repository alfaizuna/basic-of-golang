package main

import "testing"

func TestValidateEmailTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"valid standard email", "user@example.com", true},
		{"valid short email", "a@b.c", true},
		{"missing at symbol", "userexample.com", false},
		{"missing dot symbol", "user@examplecom", false},
		{"too short", "a@.", false},
		{"empty email", "", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ValidateEmail(tc.email)
			if got != tc.expected {
				t.Errorf("ValidateEmail(%q) = %t, want %t", tc.email, got, tc.expected)
			}
		})
	}
}
