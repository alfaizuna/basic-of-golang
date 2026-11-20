package main

import "testing"

func TestFakeFetcher(t *testing.T) {
	fetcher := &FakeFetcher{
		Data: map[string]*FakeResponse{
			"http://a.com": {Body: "A", URLs: []string{"http://b.com"}},
		},
	}
	var f Fetcher = fetcher

	body, urls, err := f.Fetch("http://a.com")
	if err != nil || body != "A" || len(urls) != 1 {
		t.Fatalf("unexpected fetch result: %v", err)
	}

	_, _, err = f.Fetch("http://unknown.com")
	if err == nil {
		t.Fatal("expected error on missing url, got nil")
	}
}
