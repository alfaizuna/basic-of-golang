package main

import (
	"testing"
)

type testFetcher struct{}

func (t *testFetcher) Fetch(url string) (string, []string, error) {
	if url == "http://root" {
		return "ok", []string{"http://root/a", "http://root/b"}, nil
	}
	if url == "http://root/a" {
		return "ok", nil, nil
	}
	return "", nil, nil
}

func TestRunCrawlerEndToEnd(t *testing.T) {
	report := RunCrawler("http://root", 2, &testFetcher{})

	if len(report.SuccessURLs) < 2 {
		t.Errorf("expected at least 2 success URLs, got %d (%v)", len(report.SuccessURLs), report.SuccessURLs)
	}
}
