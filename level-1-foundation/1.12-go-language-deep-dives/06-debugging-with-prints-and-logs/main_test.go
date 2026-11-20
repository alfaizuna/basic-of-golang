package main

import (
	"strings"
	"testing"
)

func TestInspectVariable(t *testing.T) {
	s := Service{Host: "db", Port: 5432}
	f, syn := InspectVariable(s)

	if !strings.Contains(f, "Host:db") || !strings.Contains(f, "Port:5432") {
		t.Errorf("fieldFormat missing field names: %q", f)
	}
	if !strings.Contains(syn, "main.Service") {
		t.Errorf("syntaxFormat should include type name: %q", syn)
	}
}
