package main

import (
	"fmt"
)

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

// TODO 1: Implementasikan method (f *Flag) Enable()
// Set f.Enabled = true (ingat cek nil guard)
func (f *Flag) Enable() {
}

// TODO 2: Implementasikan method (f *Flag) Disable()
// Set f.Enabled = false (ingat cek nil guard)
func (f *Flag) Disable() {
}

// TODO 3: Implementasikan method (f *Flag) SetRollout(percentage int) error
// - Cek nil guard
// - Validasi percentage antara 0 s.d 100
// - Set f.Rollout = percentage
func (f *Flag) SetRollout(percentage int) error {
	return nil
}

func (s *Store) FindOrCreate(name string) *Flag {
	if fl, exists := s.flags[name]; exists {
		return fl
	}
	fl := &Flag{
		Name:    name,
		Enabled: false,
		Rollout: 0,
	}
	s.flags[name] = fl
	return fl
}

func (s *Store) Get(name string) (*Flag, bool) {
	fl, ok := s.flags[name]
	return fl, ok
}

func main() {
	store := NewStore()

	// 1. Dapatkan pointer flag dari store
	flagRef := store.FindOrCreate("beta-ui")
	fmt.Printf("Initial: %s (enabled: %t, rollout: %d%%)\n", flagRef.Name, flagRef.Enabled, flagRef.Rollout)

	// 2. Modifikasi objek langsung lewat pointer
	flagRef.Enable()
	_ = flagRef.SetRollout(75)

	// 3. Ambil ulang flag dari store lewat komponen lain
	retrievedFlag, _ := store.Get("beta-ui")
	fmt.Println("After update via pointer:")
	fmt.Printf("Retrieved from store: %s (enabled: %t, rollout: %d%%)\n",
		retrievedFlag.Name, retrievedFlag.Enabled, retrievedFlag.Rollout)
	fmt.Printf("Shared memory confirmed: %t\n", flagRef == retrievedFlag)
}
