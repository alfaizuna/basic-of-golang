package main

import (
	"sync/atomic"
	"testing"
)

func TestCrawlConcurrentMock(t *testing.T) {
	urls := []string{"a", "b", "c", "d"}
	var counter int32

	CrawlConcurrentMock(urls, func(u string) {
		atomic.AddInt32(&counter, 1)
	})

	if counter != 4 {
		t.Errorf("expected 4 completed tasks, got %d", counter)
	}
}
