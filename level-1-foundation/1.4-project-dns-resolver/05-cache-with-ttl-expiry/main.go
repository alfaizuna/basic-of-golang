package main

import (
	"fmt"
	"time"
)

type RecordType string

const (
	TypeA     RecordType = "A"
	TypeCNAME RecordType = "CNAME"
)

type Record struct {
	Name  string
	Type  RecordType
	Value string
	TTL   int
}

func NewRecord(name string, recordType RecordType, value string, ttl int) *Record {
	return &Record{
		Name:  name,
		Type:  recordType,
		Value: value,
		TTL:   ttl,
	}
}

// TODO 1: Definisikan struct CacheEntry dengan field:
// - Record (*Record)
// - ExpiresAt (time.Time)
type CacheEntry struct {
	// lengkapi di sini
}

// TODO 2: Definisikan struct DNSCache dengan field:
// - entries (map[string]*CacheEntry)
// - now (func() time.Time)
type DNSCache struct {
	// lengkapi di sini
}

// TODO 3: Buat fungsi NewDNSCache() *DNSCache
func NewDNSCache() *DNSCache {
	return nil
}

// TODO 4: Buat method (c *DNSCache) Set(rec *Record)
func (c *DNSCache) Set(rec *Record) {
}

// TODO 5: Buat method (c *DNSCache) Get(name string, recordType RecordType) *Record
func (c *DNSCache) Get(name string, recordType RecordType) *Record {
	return nil
}

func main() {
	cache := NewDNSCache()
	baseTime := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	currentTime := baseTime

	// Gunakan mock clock
	cache.now = func() time.Time {
		return currentTime
	}

	rec := NewRecord("example.com", TypeA, "93.184.216.34", 300) // TTL 300 detik
	cache.Set(rec)

	// Lookup instan (masih segar)
	cached := cache.Get("example.com", TypeA)
	if cached != nil {
		fmt.Printf("Cached lookup (immediate): %s\n", cached.Value)
	}

	// Majukan waktu 301 detik (sudah kedaluwarsa)
	currentTime = baseTime.Add(301 * time.Second)

	expired := cache.Get("example.com", TypeA)
	if expired == nil {
		fmt.Println("Cached lookup (after expiry): nil")
	}
}
