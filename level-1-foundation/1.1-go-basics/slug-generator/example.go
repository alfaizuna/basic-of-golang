//go:build ignore

package main

import (
	"fmt"
	"strings"
)

// collapseSeparators mengganti spasi berulang menjadi satu strip (-)
func collapseSeparators(input string) string {
	var result string
	inSpace := false

	for _, r := range input {
		if r == ' ' {
			if !inSpace && len(result) > 0 {
				result += "-"
				inSpace = true
			}
		} else {
			result += string(r)
			inSpace = false
		}
	}
	return strings.Trim(result, "-")
}

func main() {
	sample := "  belajar   golang   dasar  "
	fmt.Println("Result:", collapseSeparators(sample)) // belajar-golang-dasar
}
