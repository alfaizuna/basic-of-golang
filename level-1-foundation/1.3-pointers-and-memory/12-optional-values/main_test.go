package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestStringPtr(t *testing.T) {
	s := "test-value"
	ptr := StringPtr(s)
	if ptr == nil {
		t.Fatalf("expected non-nil pointer from StringPtr")
	}
	if *ptr != s {
		t.Errorf("expected '%s', got '%s'", s, *ptr)
	}
}

func TestResolveServerConfigDefault(t *testing.T) {
	cfg := ServerConfig{}
	res := ResolveServerConfig(cfg)

	if res.Host != "localhost" {
		t.Errorf("expected default Host 'localhost', got '%s'", res.Host)
	}
	if res.Port != 8080 {
		t.Errorf("expected default Port 8080, got %d", res.Port)
	}
	if res.Debug != false {
		t.Errorf("expected default Debug false, got %t", res.Debug)
	}
}

func TestResolveServerConfigCustom(t *testing.T) {
	h := "192.168.1.100"
	p := 4000
	d := true

	cfg := ServerConfig{
		Host:  &h,
		Port:  &p,
		Debug: &d,
	}
	res := ResolveServerConfig(cfg)

	if res.Host != h {
		t.Errorf("expected Host '%s', got '%s'", h, res.Host)
	}
	if res.Port != p {
		t.Errorf("expected Port %d, got %d", p, res.Port)
	}
	if res.Debug != d {
		t.Errorf("expected Debug %t, got %t", d, res.Debug)
	}
}

func TestResolveServerConfigExplicitZero(t *testing.T) {
	// Memastikan nilai 0 dianggap valid jika pointer tidak nil
	p := 0
	cfg := ServerConfig{
		Port: &p,
	}
	res := ResolveServerConfig(cfg)

	if res.Port != 0 {
		t.Errorf("expected Port 0 (explicitly set), got %d", res.Port)
	}
	if res.Host != "localhost" {
		t.Errorf("expected default Host 'localhost', got '%s'", res.Host)
	}
}

func TestMainOutput(t *testing.T) {
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	output := strings.TrimSpace(buf.String())

	expectedLines := []string{
		"Default Config: localhost:8080 (debug: false)",
		"Custom Config: 0.0.0.0:3000 (debug: true)",
		"Explicit Zero Config: localhost:0 (debug: false)",
	}

	expected := strings.Join(expectedLines, "\n")
	if output != expected {
		t.Errorf("unexpected output:\nGOT:\n%s\n\nWANT:\n%s", output, expected)
	}
}
