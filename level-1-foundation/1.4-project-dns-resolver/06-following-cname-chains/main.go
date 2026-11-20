package main

import (
	"errors"
	"fmt"
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

type Zone struct {
	Name    string
	records map[string][]*Record
}

func NewZone(name string) *Zone {
	return &Zone{
		Name:    name,
		records: make(map[string][]*Record),
	}
}

func (z *Zone) AddRecord(rec *Record) {
	if z == nil || rec == nil {
		return
	}
	z.records[rec.Name] = append(z.records[rec.Name], rec)
}

func (z *Zone) Lookup(name string, recordType RecordType) *Record {
	if z == nil || z.records == nil {
		return nil
	}
	for _, r := range z.records[name] {
		if r.Type == recordType {
			return r
		}
	}
	return nil
}

// TODO: Implementasikan method (z *Zone) Resolve(name string) (*Record, error)
// Ikuti rantai CNAME berulang sampai menemukan record TypeA.
// Jika tidak ada record yang cocok, kembalikan nil, errors.New("record not found")
func (z *Zone) Resolve(name string) (*Record, error) {
	return nil, errors.New("not implemented")
}

func main() {
	z := NewZone("example.com")
	// Rantai: app -> lb -> cluster (A: 93.184.216.34)
	z.AddRecord(NewRecord("app.example.com", TypeCNAME, "lb.example.com", 300))
	z.AddRecord(NewRecord("lb.example.com", TypeCNAME, "cluster.example.com", 300))
	z.AddRecord(NewRecord("cluster.example.com", TypeA, "93.184.216.34", 300))

	rec, err := z.Resolve("app.example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Resolved app.example.com -> %s\n", rec.Value)
}
