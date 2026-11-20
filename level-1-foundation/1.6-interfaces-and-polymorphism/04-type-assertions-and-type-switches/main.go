package main

import "fmt"

// FormatData memformat nilai berdasar tipe aslinya menggunakan type switch.
func FormatData(val any) string {
	// TODO: Implementasikan type switch untuk string, int, bool, dan default
	return ""
}

func main() {
	fmt.Println(FormatData("golang"))
	fmt.Println(FormatData(10))
	fmt.Println(FormatData(true))
	fmt.Println(FormatData(false))
	fmt.Println(FormatData(3.14))
}
