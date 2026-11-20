package main

import (
	"fmt"
	"net/url"
)

// IsAllowed mengecek apakah targetURL diperbolehkan berdasarkan kedalaman dan domain.
func IsAllowed(targetURL, baseHost string, currentDepth, maxDepth int) bool {
	_ = url.Parse
	// TODO: Cek batas kedalaman dan kesamaan host domain
	return false
}

func main() {
	ok := IsAllowed("https://myblog.com/posts/1", "myblog.com", 1, 2)
	fmt.Printf("Allowed: %t\n", ok)
}
