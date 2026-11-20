package main

import (
	"fmt"
	"sync"
)

// CrawlConcurrentMock mengeksekusi processFn untuk setiap url secara paralel menggunakan goroutine dan WaitGroup.
func CrawlConcurrentMock(urls []string, processFn func(string)) {
	var _ sync.WaitGroup
	// TODO: Luncurkan goroutine untuk tiap url dan sinkronisasi dengan WaitGroup
}

func main() {
	urls := []string{"url1", "url2", "url3"}
	CrawlConcurrentMock(urls, func(u string) {
	})
	fmt.Println("All workers completed")
}
