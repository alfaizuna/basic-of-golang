package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDNSCacheExpiry(t *testing.T) {
	cache := NewDNSCache()
	if cache == nil || cache.entries == nil || cache.now == nil {
		t.Fatal("NewDNSCache did not initialize fields properly")
	}

	mockTime := time.Date(2026, 6, 1, 10, 0, 0, 0, time.UTC)
	cache.now = func() time.Time {
		return mockTime
	}

	rec := NewRecord("golang.org", TypeA, "142.250.190.49", 60)
	cache.Set(rec)

	// Harusnya ditemukan sebelum 60 detik
	mockTime = mockTime.Add(30 * time.Second)
	got := cache.Get("golang.org", TypeA)
	if got != rec {
		t.Errorf("expected to retrieve cached record before TTL expires")
	}

	// Setelah lewat 60 detik
	mockTime = mockTime.Add(31 * time.Second)
	expired := cache.Get("golang.org", TypeA)
	if expired != nil {
		t.Errorf("expected nil after TTL expiry, got %v", expired)
	}

	// Nil safety checks
	var nilCache *DNSCache
	nilCache.Set(rec)
	if r := nilCache.Get("golang.org", TypeA); r != nil {
		t.Errorf("expected nil from nil cache")
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
	want1 := "Cached lookup (immediate): 93.184.216.34"
	want2 := "Cached lookup (after expiry): nil"

	if !strings.Contains(got, want1) {
		t.Errorf("output does not contain %q, got:\n%s", want1, got)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("output does not contain %q, got:\n%s", want2, got)
	}
}
