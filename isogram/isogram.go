package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram checks whether a given string is an isogram
func IsIsogram(input string) bool {
	included := map[rune]bool{}
	for _, c := range strings.ToLower(input) {
		if !unicode.IsLetter(c) {
			continue
		}
		if included[c] {
			return false
		}
		included[c] = true
	}
	return true
}
