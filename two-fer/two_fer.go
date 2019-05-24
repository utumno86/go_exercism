package twofer

import "fmt"

// ShareWith function returns two for me and one for you
func ShareWith(input string) string {
	if input != "" {
		return fmt.Sprintf("One for %s, one for me.", input)
	}
	return "One for you, one for me."
}