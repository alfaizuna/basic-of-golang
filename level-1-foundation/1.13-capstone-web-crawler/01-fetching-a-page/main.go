package main

import (
	"fmt"
)

// Fetcher mendefinisikan antarmuka pengambil data URL.
type Fetcher interface {
	Fetch(url string) (string, []string, error)
}

type FakeResponse struct {
	Body string
	URLs []string
}

type FakeFetcher struct {
	Data map[string]*FakeResponse
}

// Fetch mengembalikan Body dan daftar URLs dari mock data.
func (f *FakeFetcher) Fetch(url string) (string, []string, error) {
	// TODO: Kembalikan Body, URLs, nil jika ada di f.Data, atau error "not found: <url>" jika tidak ditemukan
	return "", nil, nil
}

func main() {
	fetcher := &FakeFetcher{
		Data: map[string]*FakeResponse{
			"https://shop.local": {
				Body: "Home Page",
				URLs: []string{"https://shop.local/products", "https://shop.local/about"},
			},
		},
	}

	body, urls, _ := fetcher.Fetch("https://shop.local")
	fmt.Printf("Fetched: %s with %d links\n", body, len(urls))
}
