package main

import "fmt"

// SafeRun mengeksekusi fn dan mengonversi panic menjadi error biasa.
func SafeRun(fn func()) (err error) {
	// TODO: Pasang defer dengan recover() untuk menangkap panic
	return nil
}

func main() {
	err := SafeRun(func() {
		panic("something went wrong")
	})
	if err != nil {
		fmt.Printf("Caught panic: %v\n", err)
	}
}
