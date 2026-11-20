//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

func crawlStub(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Crawling:", url)
	time.Sleep(10 * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup
	urls := []string{"https://site.com/a", "https://site.com/b"}

	for _, u := range urls {
		wg.Add(1)
		go crawlStub(u, &wg)
	}

	wg.Wait()
	fmt.Println("Selesai")
}
