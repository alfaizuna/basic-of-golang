package main

import "fmt"

// TODO 1: Definisikan struct Flag dengan field:
// - Name (string)
// - Enabled (bool)
// - Rollout (int)

type Flag struct {
	// lengkapi field di sini
}

// TODO 2: Definisikan struct Store dengan field:
// - flags (map[string]*Flag)

type Store struct {
	// lengkapi field di sini
}

// TODO 3: Buat fungsi konstruktor NewStore() *Store
func NewStore() *Store {
	return nil
}

// TODO 4: Buat method (s *Store) Count() int
func (s *Store) Count() int {
	return 0
}

func main() {
	store := NewStore()
	if store != nil {
		fmt.Printf("Store created. Initial flag count: %d\n", store.Count())
	}
}
