package main

import "fmt"

// FormatModuleHeader membuat representasi header file go.mod.
func FormatModuleHeader(modulePath, goVersion string) string {
	// TODO: Kembalikan format string "module <modulePath>\n\ngo <goVersion>"
	return ""
}

func main() {
	header := FormatModuleHeader("github.com/user/myapp", "1.23")
	fmt.Println(header)
}
