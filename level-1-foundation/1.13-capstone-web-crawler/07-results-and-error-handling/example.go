//go:build ignore

package main

import "fmt"

type CrawlReport struct {
	SuccessURLs []string
	FailedURLs  map[string]string
}

func main() {
	fmt.Println("Crawler capstone wrap-up")
}
