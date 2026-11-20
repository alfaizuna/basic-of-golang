package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (string, []string, error)
}

type CrawlReport struct {
	mu          sync.Mutex
	SuccessURLs []string
	FailedURLs  map[string]string
}

func NewCrawlReport() *CrawlReport {
	return &CrawlReport{
		FailedURLs: make(map[string]string),
	}
}

// AddSuccess menambahkan url sukses ke report secara thread-safe.
func (r *CrawlReport) AddSuccess(url string) {
	// TODO: Lindungi dengan mutex dan append ke r.SuccessURLs
}

// AddFailure mencatat error pada url ke report secara thread-safe.
func (r *CrawlReport) AddFailure(url string, errMsg string) {
	// TODO: Lindungi dengan mutex dan setel r.FailedURLs[url] = errMsg
}

// RunCrawler mengeksekusi crawl pohon URL hingga maxDepth secara konkuren.
func RunCrawler(startURL string, maxDepth int, fetcher Fetcher) *CrawlReport {
	// TODO: Implementasikan crawling konkuren menggunakan goroutine, WaitGroup, dan tracker
	return NewCrawlReport()
}

func main() {
	fmt.Println("Capstone crawler wrap-up ready")
}
