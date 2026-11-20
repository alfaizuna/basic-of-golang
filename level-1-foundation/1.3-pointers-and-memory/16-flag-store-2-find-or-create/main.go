package main

import "fmt"

type Flag struct {
	Name    string
	Enabled bool
	Rollout int
}

type Store struct {
	flags map[string]*Flag
}

func NewStore() *Store {
	return &Store{
		flags: make(map[string]*Flag),
	}
}

func (s *Store) Count() int {
	if s == nil || s.flags == nil {
		return 0
	}
	return len(s.flags)
}

// Get mencari flag berdasarkan nama.
// Kembalikan (*Flag, true) jika ditemukan, atau (nil, false) jika tidak ada.
func (s *Store) Get(name string) (*Flag, bool) {
	// TODO: cari flag di s.flags
	return nil, false
}

// FindOrCreate mencari flag berdasarkan nama.
// Jika sudah ada, kembalikan pointer flag tersebut.
// Jika belum ada, buat flag baru dengan default (Enabled: false, Rollout: 0),
// simpan ke s.flags, lalu kembalikan pointernya.
func (s *Store) FindOrCreate(name string) *Flag {
	// TODO: implementasikan logika find-or-create
	return nil
}

func main() {
	store := NewStore()

	// 1. Buat flag baru lewat FindOrCreate
	f1 := store.FindOrCreate("dark-mode")
	fmt.Printf("Flag count: %d\n", store.Count())
	fmt.Printf("Flag 1: %s (enabled: %t, rollout: %d%%)\n", f1.Name, f1.Enabled, f1.Rollout)

	// 2. Panggil lagi dengan nama yang sama -> harus mengembalikan pointer yang sama
	f2 := store.FindOrCreate("dark-mode")
	fmt.Printf("Same memory address: %t\n", f1 == f2)
}
