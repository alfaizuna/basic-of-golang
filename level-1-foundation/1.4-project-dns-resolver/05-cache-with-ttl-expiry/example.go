//go:build ignore

package main

import "time"

// example.go — hanya untuk referensi, jangan diubah

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

type CacheEntry struct {
	Record    *Record
	ExpiresAt time.Time
}

type DNSCache struct {
	entries map[string]*CacheEntry
	now     func() time.Time
}

func NewDNSCache() *DNSCache {
	return &DNSCache{
		entries: make(map[string]*CacheEntry),
		now:     time.Now,
	}
}
