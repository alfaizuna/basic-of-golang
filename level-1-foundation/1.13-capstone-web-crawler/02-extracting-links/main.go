package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// ExtractLinks mengekstrak semua hyperlink dan mengubahnya menjadi absolute URL.
func ExtractLinks(baseStr, htmlContent string) ([]string, error) {
	_ = url.Parse
	_ = strings.NewReader
	_ = html.NewTokenizer
	// TODO: Parse baseURL, gunakan html.NewTokenizer, dan resolve href atribut
	return nil, nil
}

func main() {
	htmlData := `<p>Kunjungi <a href="/products">Katalog</a> kami.</p>`
	links, _ := ExtractLinks("https://store.com/", htmlData)
	for _, l := range links {
		fmt.Println("Extracted link:", l)
	}
}
