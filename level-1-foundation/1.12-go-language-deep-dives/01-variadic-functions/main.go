package main

import (
	"fmt"
	"strings"
)

// JoinStrings menggabungkan argumen variadik words dengan pemisah sep.
func JoinStrings(sep string, words ...string) string {
	_ = strings.Join
	// TODO: Gabungkan words dengan strings.Join
	return ""
}

func main() {
	res := JoinStrings("-", "go", "is", "fast")
	fmt.Println(res)
}
