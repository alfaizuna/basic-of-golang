package main

import "fmt"

// TODO 1: Definisikan tipe RecordType berbasis string
// type RecordType string

// TODO 2: Definisikan konstanta TypeA ("A") dan TypeCNAME ("CNAME")
// const (
//     TypeA     RecordType = "A"
//     TypeCNAME RecordType = "CNAME"
// )

// TODO 3: Definisikan struct Record dengan field:
// - Name (string)
// - Type (RecordType)
// - Value (string)
// - TTL (int)

// TODO 4: Buat fungsi NewRecord(name string, recordType RecordType, value string, ttl int) *Record
// func NewRecord(name string, recordType RecordType, value string, ttl int) *Record { ... }

func main() {
	// rec := NewRecord("example.com", TypeA, "93.184.216.34", 300)
	// cname := NewRecord("www.example.com", TypeCNAME, "example.com", 3600)
	// fmt.Printf("Record created: %s (%s -> %s, TTL: %ds)\n", rec.Name, rec.Type, rec.Value, rec.TTL)
	// fmt.Printf("Alias created: %s (%s -> %s, TTL: %ds)\n", cname.Name, cname.Type, cname.Value, cname.TTL)

	fmt.Println("Record model placeholder")
}
