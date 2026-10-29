package main

import (
	"fmt"
	"strings"
	"errors"
)

// isAllowedSlugChar memeriksa apakah rune adalah alfanumerik, spasi, atau dash (-).
func isAllowedSlugChar(r rune) bool {
	isAlphaNumeric := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
	isSep := r == ' ' || r == '-'
	return isAlphaNumeric || isSep
}

// countInvalidChars menghitung banyaknya karakter yang tidak diizinkan pada text.
func countInvalidChars(text string) int {
	count := 0
	for _, r := range text {
		if !isAllowedSlugChar(r) {
			count++
		}
	}
	return count
}

func cleanText(text string) string {
	var filtered string
	for _, r := range text {
		if isAllowedSlugChar(r) {
			filtered += string(r)
		}
	}
	return strings.TrimSpace(strings.ToLower(filtered))
}

func transformToSlug(text string) string {
	cleaned := cleanText(text)
	var result string
	inSeparator := false
	for _, r := range cleaned {
		if r == ' ' || r == '-' {
			if !inSeparator && len(result) > 0 {
				result += "-"
				inSeparator = true 
			} 
		} else {
				result += string(r) 
				inSeparator = false
		}
	}
	return strings.Trim(result, "-")
}

func GenerateSlug(title string) (string, error) {
	if strings.TrimSpace(title) == "" {
		return "", errors.New("input text cannot be empty")
	}
	slug := transformToSlug(title)
	if slug == "" {
		return "", errors.New("no valid slug characters found")
	}
	return slug, nil
}

func main() {
	title1 := "  Belajar Go: Dasar #1 (Pemula)!  "
	slug1, err1 := GenerateSlug(title1)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Printf("Input: %q\nSlug:  %q\n", title1, slug1)
	}

	title2 := "### ??? !!!"
	_, err2 := GenerateSlug(title2)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}

}
