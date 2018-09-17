package scrabble

import "unicode"

// Score from a set of letter in Scrabble
func Score(input string) int {
	scoreTotal := 0

	for _, char := range input {
		switch unicode.ToUpper(char) {
		case 'D', 'G':
			scoreTotal += 2
		case 'B', 'C', 'M', 'P':
			scoreTotal += 3
		case 'F', 'H', 'V', 'W', 'Y':
			scoreTotal += 4
		case 'K':
			scoreTotal += 5
		case 'J', 'X':
			scoreTotal += 8
		case 'Q', 'Z':
			scoreTotal += 10
		default:
			scoreTotal++
		}
	}

	return scoreTotal
}
