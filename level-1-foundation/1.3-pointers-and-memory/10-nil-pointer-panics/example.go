package main

import "fmt"

// example.go — hanya untuk referensi, jangan diubah

type LogEntry struct {
	Level   string
	Message string
}

// formatLog menerima pointer ke LogEntry.
// Jika entry adalah nil, mengakses entry.Level akan memicu panic!
func formatLog(entry *LogEntry) string {
	if entry == nil {
		return "[NIL LOG ENTRY]"
	}
	return fmt.Sprintf("[%s] %s", entry.Level, entry.Message)
}

var _ = fmt.Sprintf // pastikan import fmt tidak unused
