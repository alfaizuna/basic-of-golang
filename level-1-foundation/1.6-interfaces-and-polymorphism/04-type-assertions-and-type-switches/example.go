//go:build ignore

package main

import "fmt"

func inspect(val any) {
	// Type switch
	switch v := val.(type) {
	case string:
		fmt.Printf("String dengan panjang %d: %q\n", len(v), v)
	case int:
		fmt.Printf("Integer kuadrat: %d\n", v*v)
	case bool:
		fmt.Printf("Boolean kebalikan: %t\n", !v)
	default:
		fmt.Printf("Tipe tidak diketahui: %T (%v)\n", v, v)
	}
}

func main() {
	// Type assertion
	var x any = "gopher"
	s, ok := x.(string)
	if ok {
		fmt.Println("Type assertion sukses:", s)
	}

	inspect(10)
	inspect("hello")
	inspect(true)
	inspect(3.14)
}
