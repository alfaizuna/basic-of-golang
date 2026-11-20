package main

import (
	"fmt"
	"testing"
)

type testFakeFetcher struct {
	Data map[string][]string
}

func (f *testFakeFetcher) Fetch(url string) (string, []string, error) {
	if urls, ok := f.Data[url]; ok {
		return "body", urls, nil
	}
	return "", nil, fmt.Errorf("not found")
}

func TestCrawlSequential(t *testing.T) {
	f := &testFakeFetcher{
		Data: map[string][]string{
			"A": {"B", "C"},
			"B": {"A"}, // cycle
			"C": {"D"},
			"D": {},
		},
	}
	visited := make(map[string]bool)
	pages := CrawlSequential("A", f, visited)

	if len(pages) != 4 {
		t.Fatalf("expected 4 unique pages, got %d (%v)", len(pages), pages)
	}
}
