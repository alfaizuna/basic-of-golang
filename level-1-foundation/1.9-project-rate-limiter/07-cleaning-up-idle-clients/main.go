package main

import (
	"fmt"
	"time"
)

type ClientInfo struct {
	lastSeen time.Time
}

type LimiterWithCleanup struct {
	clients map[string]*ClientInfo
	now     func() time.Time
}

func NewLimiterWithCleanup(now func() time.Time) *LimiterWithCleanup {
	if now == nil {
		now = time.Now
	}
	return &LimiterWithCleanup{
		clients: make(map[string]*ClientInfo),
		now:     now,
	}
}

// Touch mencatat aktivitas terakhir klien pada waktu saat ini.
func (l *LimiterWithCleanup) Touch(key string) {
	// TODO: Simpan entri baru di l.clients dengan lastSeen = l.now()
}

// Cleanup menghapus klien yang waktu idle-nya melebihi maxIdle.
func (l *LimiterWithCleanup) Cleanup(maxIdle time.Duration) int {
	// TODO: Iterasi l.clients, hapus jika cur.Sub(info.lastSeen) > maxIdle, dan kembalikan total yang dihapus
	return 0
}

func main() {
	curr := time.Now()
	l := NewLimiterWithCleanup(func() time.Time { return curr })

	l.Touch("client-active")
	l.Touch("client-idle")

	curr = curr.Add(10 * time.Minute)
	l.Touch("client-active")

	removed := l.Cleanup(5 * time.Minute)
	fmt.Printf("Removed clients: %d\n", removed)
	fmt.Printf("Remaining clients: %d\n", len(l.clients))
}
