package main

import (
	"strings"
	"testing"
)

func TestToJSON(t *testing.T) {
	u := User{ID: 10, Username: "alice"}
	s, err := ToJSON(u)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(s, `"id":10`) || !strings.Contains(s, `"username":"alice"`) {
		t.Errorf("missing fields in json: %s", s)
	}
	if strings.Contains(s, "email") {
		t.Errorf("empty email should be omitted by omitempty, got: %s", s)
	}
}
