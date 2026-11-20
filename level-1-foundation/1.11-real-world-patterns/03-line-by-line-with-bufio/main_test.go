package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestFilterLines(t *testing.T) {
	input := "apple\nbanana\npineapple\norange"
	got, err := FilterLines(strings.NewReader(input), "apple")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []string{"apple", "pineapple"}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
