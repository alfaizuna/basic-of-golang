package main

import "fmt"

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

// TODO 1: Definisikan struct Zone dengan field:
// - Name (string)
// - records (map[string][]*Record)
type Zone struct {
	// lengkapi di sini
}

// TODO 2: Buat fungsi NewZone(name string) *Zone
func NewZone(name string) *Zone {
	return nil
}

// TODO 3: Buat method (z *Zone) RecordCount() int
func (z *Zone) RecordCount() int {
	return 0
}

func main() {
	z := NewZone("example.com")
	if z != nil {
		fmt.Printf("Zone created: %s (records: %d)\n", z.Name, z.RecordCount())
	}
}
