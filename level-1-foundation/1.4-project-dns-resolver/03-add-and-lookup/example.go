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
