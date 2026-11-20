//go:build ignore

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	s := strings.TrimSpace("   42   ")
	n, _ := strconv.Atoi(s)
	fmt.Printf("Angka: %d, Tanggal: %s\n", n, time.Now().Format("2006-01-02"))
}
