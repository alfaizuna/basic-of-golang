package main

import "fmt"

// 1. Buat konstanta MaxScore bernilai 100
const MaxScore int = 100

// 2. Buat blok const (...) dengan iota untuk StatusDraft dan StatusPublished
const (
	StatusDraft = iota
	StatusPublished
	StatusAJa
)
func main() {
	// 3. Deklarasikan variabel 'defaultAttempts' bertipe int (tanpa nilai awal)
	var defaultAttempts int 
	// 4. Deklarasikan variabel 'isReady' bertipe bool (tanpa nilai awal)
	var isReady bool
	// 5. Cetak MaxScore, StatusDraft, StatusPublished, defaultAttempts, dan isReady
	fmt.Println(MaxScore)
	fmt.Println(StatusDraft)
	fmt.Println(StatusPublished)
	fmt.Println(StatusAJa)
	fmt.Println(defaultAttempts)
	fmt.Println(isReady)

}
