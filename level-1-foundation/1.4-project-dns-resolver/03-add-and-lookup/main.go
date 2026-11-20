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

// TODO 1: Implementasikan method (z *Zone) AddRecord(rec *Record)
func (z *Zone) AddRecord(rec *Record) {
}

// TODO 2: Implementasikan method (z *Zone) Lookup(name string, recordType RecordType) *Record
func (z *Zone) Lookup(name string, recordType RecordType) *Record {
	return nil
}

func main() {
	z := NewZone("example.com")
	z.AddRecord(NewRecord("example.com", TypeA, "93.184.216.34", 300))

	found := z.Lookup("example.com", TypeA)
	if found != nil {
		fmt.Printf("Lookup example.com (A): %s\n", found.Value)
	}

	missing := z.Lookup("notfound.com", TypeA)
	if missing == nil {
		fmt.Println("Lookup notfound.com (A): nil")
	}
}
