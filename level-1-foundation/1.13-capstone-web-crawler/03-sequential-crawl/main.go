package main

import (
	"fmt"
)

type Fetcher interface {
	Fetch(url string) (string, []string, error)
}

// CrawlSequential melakukan penelusuran sekuensial dengan visited set.
func CrawlSequential(url string, fetcher Fetcher, visited map[string]bool) []string {
	// TODO: Cek visited, tandai kunjungan, ambil link anak, dan panggil rekursif
	return nil
}

func main() {
	visited := make(map[string]bool)
	_ = visited
	fmt.Println("Sequential crawler skeleton")
}
