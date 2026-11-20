//go:build ignore

package main

import (
	"fmt"
	"net/url"
	"strings"
	"golang.org/x/net/html"
)

func main() {
	baseURL, _ := url.Parse("https://example.com/blog/")
	relURL, _ := url.Parse("../about")
	resolved := baseURL.ResolveReference(relURL)
	fmt.Println("Absolut URL:", resolved.String())
}
