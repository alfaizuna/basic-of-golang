//go:build ignore

package main

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

func (z *Zone) RecordCount() int {
	if z == nil || z.records == nil {
		return 0
	}
	total := 0
	for _, recs := range z.records {
		total += len(recs)
	}
	return total
}
