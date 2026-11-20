//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[string]time.Time)
	m["clientA"] = time.Now().Add(-10 * time.Minute)
	// delete key dari map
	delete(m, "clientA")
	fmt.Println("Remaining:", len(m))
}
