package main

import "fmt"

// FormatSemver memformat versi semantik sesuai standar Go Modules.
func FormatSemver(major, minor, patch int, prerelease string) string {
	// TODO: Kembalikan string semver
	return ""
}

func main() {
	fmt.Println("Release:", FormatSemver(1, 2, 3, ""))
	fmt.Println("Pre-release:", FormatSemver(2, 0, 0, "beta.1"))
}
