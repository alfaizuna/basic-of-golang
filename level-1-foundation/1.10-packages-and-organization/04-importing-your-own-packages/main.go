package main

import "fmt"

// ResolveNamespace mendeskripsikan strategi resolusi namespace import.
func ResolveNamespace(alias, originalPackage string) string {
	// TODO: Kembalikan string sesuai keberadaan alias
	return ""
}

func main() {
	fmt.Println(ResolveNamespace("crand", "crypto/rand"))
	fmt.Println(ResolveNamespace("", "strings"))
}
