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

func (z *Zone) AddRecord(rec *Record) {
	if z == nil || rec == nil {
		return
	}
	z.records[rec.Name] = append(z.records[rec.Name], rec)
}

func (z *Zone) Transfer() *Zone {
	if z == nil {
		return nil
	}
	target := NewZone(z.Name)
	for _, recs := range z.records {
		for _, r := range recs {
			copiedRec := &Record{
				Name:  r.Name,
				Type:  r.Type,
				Value: r.Value,
				TTL:   r.TTL,
			}
			target.AddRecord(copiedRec)
		}
	}
	return target
}
