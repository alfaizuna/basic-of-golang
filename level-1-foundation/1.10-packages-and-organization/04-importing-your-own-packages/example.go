//go:build ignore

package main

import (
	"fmt"
	crand "crypto/rand"
	mrand "math/rand"
)

func main() {
	fmt.Println("Alias import mencegah benturan nama antar package")
	_ = crand.Reader
	_ = mrand.Int()
}
