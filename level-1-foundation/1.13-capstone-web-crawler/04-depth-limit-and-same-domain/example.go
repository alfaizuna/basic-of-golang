//go:build ignore

package main

import (
	"fmt"
	"net/url"
)

func isSameDomain(targetStr, allowedHost string) bool {
	u, err := url.Parse(targetStr)
	if err != nil {
		return false
	}
	return u.Host == allowedHost
}

func main() {
	fmt.Println("Same domain check:", isSameDomain("https://shop.com/deals", "shop.com"))
}
