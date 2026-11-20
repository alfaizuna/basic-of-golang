package main

import "fmt"

// ValidateImportCycle mendeteksi apakah terdapat import cycle pada graf dependensi package.
func ValidateImportCycle(dependencies map[string][]string, start string) bool {
	// TODO: Lakukan penelusuran DFS untuk mendeteksi siklus dependensi
	return false
}

func main() {
	deps := map[string][]string{
		"A": {"B"},
		"B": {"C"},
		"C": {"A"},
	}
	fmt.Println("Has cycle:", ValidateImportCycle(deps, "A"))
}
