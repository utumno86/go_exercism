package luhn

import (
	"strings"
)

// Valid checks whether a series of numbers are valid as credit card numbers
func Valid(input string) bool {
	strippedString := strings.Replace(input, " ", "", -1)
	sum := checksum(strippedString)
	if len(strippedString) < 2 {
		return false
	}
	return (sum % 10) == 0
}

func checksum(input string) int {
	evenNums := len(input)%2 == 0
	sum := int(0)
	for i, num := range []rune(input) {
		if num < 48 || num > 57 {
			return -1
		}
		currentDigit := int(num - '0')
		evenOddIndex := i
		if !evenNums {
			evenOddIndex++
		}
		sum = sum + currentDigit
		if (evenOddIndex % 2) == 0 {
			sum = sum + currentDigit
			if currentDigit*2 > 9 {
				sum = sum - 9
			}
		}
	}
	return sum
}
