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

func (r *Record) UpdateValue(newValue string) {
	if r == nil {
		return
	}
	r.Value = newValue
}

func (r *Record) UpdateTTL(newTTL int) {
	if r == nil {
		return
	}
	r.TTL = newTTL
}
