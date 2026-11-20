//go:build ignore

package main

import (
	"fmt"
)

type FetchResult struct {
	Body string
	URLs []string
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func main() {
	fmt.Println("Fetcher interface model")
}
