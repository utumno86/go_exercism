package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram checks whether a given string is an isogram
func IsIsogram(input string) bool {
	var included []rune
	for _, c := range strings.ToLower(input) {
		if contains(included, c) {
			return false
		}
		included = append(included, c)
	}
	return true
}

func contains(included []rune, c rune) bool {
	for _, a := range included {
		if a == c && unicode.IsLetter(c) {
			return true
		}
	}
	return false
}
