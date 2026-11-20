//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	raw := "Baris 1\nBaris 2\nBaris 3"
	scanner := bufio.NewScanner(strings.NewReader(raw))

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Dibaca:", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error saat scanning:", err)
	}
}
