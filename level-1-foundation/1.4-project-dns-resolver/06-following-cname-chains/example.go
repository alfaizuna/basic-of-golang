//go:build ignore

package main

import "errors"

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

func (z *Zone) Resolve(name string) (*Record, error) {
	current := name
	for {
		if rec := z.Lookup(current, TypeA); rec != nil {
			return rec, nil
		}

		cname := z.Lookup(current, TypeCNAME)
		if cname == nil {
			return nil, errors.New("record not found")
		}

		current = cname.Value
	}
}
